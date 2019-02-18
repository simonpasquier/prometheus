// Copyright 2016 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package treecache

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/samuel/go-zookeeper/zk"
)

var (
	failureCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "prometheus",
		Subsystem: "treecache",
		Name:      "zookeeper_failures_total",
		Help:      "The total number of ZooKeeper failures.",
	})
	numWatchers = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "prometheus",
		Subsystem: "treecache",
		Name:      "watcher_goroutines",
		Help:      "The current number of watcher goroutines.",
	})
)

func init() {
	prometheus.MustRegister(failureCounter)
	prometheus.MustRegister(numWatchers)
}

// ZookeeperLogger wraps a log.Logger into a zk.Logger.
type ZookeeperLogger struct {
	logger log.Logger
}

// NewZookeeperLogger is a constructor for ZookeeperLogger.
func NewZookeeperLogger(logger log.Logger) ZookeeperLogger {
	return ZookeeperLogger{logger: logger}
}

// Printf implements zk.Logger.
func (zl ZookeeperLogger) Printf(s string, i ...interface{}) {
	level.Info(zl.logger).Log("msg", fmt.Sprintf(s, i...))
}

// A ZookeeperTreeCache keeps data from all children of a Zookeeper path
// locally cached and updated according to received events.
type ZookeeperTreeCache struct {
	conn   *zk.Conn
	prefix string
	events chan ZookeeperTreeCacheEvent
	ctx    context.Context
	head   *zookeeperTreeCacheNode

	logger log.Logger
}

// A ZookeeperTreeCacheEvent models a Zookeeper event for a path.
type ZookeeperTreeCacheEvent struct {
	Path string
	Data *[]byte
}

type zookeeperTreeCacheNode struct {
	data     *[]byte
	events   chan zk.Event
	children map[string]*zookeeperTreeCacheNode
}

// NewZookeeperTreeCache creates a new ZookeeperTreeCache for a given path.
func NewZookeeperTreeCache(ctx context.Context, conn *zk.Conn, path string, events chan ZookeeperTreeCacheEvent, logger log.Logger) *ZookeeperTreeCache {
	tc := &ZookeeperTreeCache{
		conn:   conn,
		prefix: path,
		events: events,
		ctx:    ctx,

		logger: logger,
	}
	tc.head = &zookeeperTreeCacheNode{
		events:   make(chan zk.Event),
		children: map[string]*zookeeperTreeCacheNode{},
	}
	go tc.loop(path)
	return tc
}

func (tc *ZookeeperTreeCache) loop(path string) {
	defer func() {
		tc.conn.Close()
	}()
	failureMode := false
	retryChan := make(chan struct{})

	failure := func() {
		failureCounter.Inc()
		failureMode = true
		select {
		case <-tc.ctx.Done():
			return
		default:
		}
		time.AfterFunc(time.Second*10, func() {
			select {
			case <-tc.ctx.Done():
				return
			default:
			}
			retryChan <- struct{}{}
		})
	}

	err := tc.recursiveNodeUpdate(path, tc.head)
	if err != nil {
		level.Error(tc.logger).Log("msg", "Error during initial read of Zookeeper", "err", err)
		failure()
	}

	for {
		select {
		case <-tc.ctx.Done():
			return
		default:
		}
		select {
		case ev := <-tc.head.events:
			level.Debug(tc.logger).Log("msg", "Received Zookeeper event", "event", ev)
			if failureMode {
				continue
			}

			if ev.Type == zk.EventNotWatching {
				level.Info(tc.logger).Log("msg", "Lost connection to Zookeeper.")
				failure()
				continue
			}
			path := strings.TrimPrefix(ev.Path, tc.prefix)
			parts := strings.Split(path, "/")
			node := tc.head
			for _, part := range parts[1:] {
				childNode := node.children[part]
				if childNode == nil {
					childNode = &zookeeperTreeCacheNode{
						events:   tc.head.events,
						children: map[string]*zookeeperTreeCacheNode{},
					}
					node.children[part] = childNode
				}
				node = childNode
			}

			err := tc.recursiveNodeUpdate(ev.Path, node)
			if err != nil {
				level.Error(tc.logger).Log("msg", "Error during processing of Zookeeper event", "err", err)
				failure()
			} else if tc.head.data == nil {
				level.Error(tc.logger).Log("msg", "Error during processing of Zookeeper event", "err", "path no longer exists", "path", tc.prefix)
				failure()
			}
		case <-retryChan:
			level.Info(tc.logger).Log("msg", "Attempting to resync state with Zookeeper")
			previousState := &zookeeperTreeCacheNode{
				children: tc.head.children,
			}
			// Reset root child nodes before traversing the Zookeeper path.
			tc.head.children = make(map[string]*zookeeperTreeCacheNode)

			if err := tc.recursiveNodeUpdate(tc.prefix, tc.head); err != nil {
				level.Error(tc.logger).Log("msg", "Error during Zookeeper resync", "err", err)
				// Revert to our previous state.
				tc.head.children = previousState.children
				failure()
				continue
			}
			tc.resyncState(tc.prefix, tc.head, previousState)
			level.Info(tc.logger).Log("Zookeeper resync successful")
			failureMode = false
		case <-tc.ctx.Done():
			return
		}
	}
}

func (tc *ZookeeperTreeCache) recursiveNodeUpdate(path string, node *zookeeperTreeCacheNode) error {
	data, _, dataWatcher, err := tc.conn.GetW(path)
	if err == zk.ErrNoNode {
		tc.recursiveDelete(path, node)
		if node == tc.head {
			return fmt.Errorf("path %s does not exist", path)
		}
		return nil
	} else if err != nil {
		return err
	}

	if node.data == nil || !bytes.Equal(*node.data, data) {
		node.data = &data
		select {
		case tc.events <- ZookeeperTreeCacheEvent{Path: path, Data: node.data}:
		case <-tc.ctx.Done():
			return tc.ctx.Err()
		}
	}

	children, _, childWatcher, err := tc.conn.ChildrenW(path)
	if err == zk.ErrNoNode {
		tc.recursiveDelete(path, node)
		return nil
	} else if err != nil {
		return err
	}

	currentChildren := map[string]struct{}{}
	for _, child := range children {
		currentChildren[child] = struct{}{}
		childNode := node.children[child]
		// The child node didn't already exist.
		if childNode == nil {
			node.children[child] = &zookeeperTreeCacheNode{
				events:   node.events,
				children: map[string]*zookeeperTreeCacheNode{},
			}
			err = tc.recursiveNodeUpdate(path+"/"+child, node.children[child])
			if err != nil {
				return err
			}
		}
	}

	// Remove nodes that no longer exist.
	for name, childNode := range node.children {
		if _, ok := currentChildren[name]; !ok || node.data == nil {
			tc.recursiveDelete(path+"/"+name, childNode)
			delete(node.children, name)
		}
	}

	go func() {
		numWatchers.Inc()
		// Pass up zookeeper events, until the node is deleted.
		select {
		case event := <-dataWatcher:
			node.events <- event
		case event := <-childWatcher:
			node.events <- event
		case <-tc.ctx.Done():
		}
		numWatchers.Dec()
	}()
	return nil
}

func (tc *ZookeeperTreeCache) resyncState(path string, currentState, previousState *zookeeperTreeCacheNode) {
	for child, previousNode := range previousState.children {
		if currentNode, present := currentState.children[child]; present {
			tc.resyncState(path+"/"+child, currentNode, previousNode)
		} else {
			tc.recursiveDelete(path+"/"+child, previousNode)
		}
	}
}

func (tc *ZookeeperTreeCache) recursiveDelete(path string, node *zookeeperTreeCacheNode) {
	if node.data != nil {
		select {
		case tc.events <- ZookeeperTreeCacheEvent{Path: path, Data: nil}:
		case <-tc.ctx.Done():
			return
		}
		node.data = nil
	}
	for name, childNode := range node.children {
		tc.recursiveDelete(path+"/"+name, childNode)
	}
}
