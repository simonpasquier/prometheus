// Copyright 2018 The Prometheus Authors
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

// Package tracing exposes a single API to manipulate OpenTracing spans along
// with user-annotated traces from the runtime/trace package.
// When the application is instrumented with the opentracing.NoopTracer, it is
// still possible to collect tracing information from the runtime.
package tracing

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"
)

type contextKey struct{}

var activeSpanKey = contextKey{}

// Span represents a trace which combines an OpenTracing span with runtime/trace user-annotations.
type Span struct {
	otSpan    opentracing.Span
	finishers []func()
}

type runtimeTracer interface {
	End()
}

type noopTracer struct{}

func (t noopTracer) End() {}

// NewSpan creates a new Span not attached to an existing OpenTracing span.
func NewSpan(ctx context.Context, name string) (*Span, context.Context) {
	return newTaskSpan(ctx, name, nil)
}

// Attach associates an existing OpenTracing span.
// The OpenTracing span won't be terminated when calling Finish().
func (s *Span) Attach(otSpan opentracing.Span) {
	if s.otSpan != nil {
		panic("OpenTracing span already attached!")
	}
	s.otSpan = otSpan
}

func newTaskSpan(ctx context.Context, name string, otSpan opentracing.Span) (*Span, context.Context) {
	sp := &Span{otSpan: otSpan, finishers: make([]func(), 0)}
	nctx, t := newTask(ctx, name)
	if t != nil {
		sp.finishers = append(sp.finishers, t.End)
	}
	if otSpan != nil {
		sp.finishers = append(sp.finishers, otSpan.Finish)
	}
	return sp, context.WithValue(nctx, activeSpanKey, sp)
}

func newRegionSpan(ctx context.Context, name string, otSpan opentracing.Span) (*Span, context.Context) {
	sp := &Span{otSpan: otSpan, finishers: []func(){otSpan.Finish}}
	r := startRegion(ctx, name)
	if r != nil {
		sp.finishers = append(sp.finishers, r.End)
	}
	return sp, context.WithValue(ctx, activeSpanKey, sp)
}

// StartSpanFromContext creates a new OpenTracing span (linked to a parent OpenTracing
// span if one is found in the current context) and starts a new task (for top-level
// spans) or a new region.
// The OpenTracing span as well as the task or region will be terminated when calling Finish().
func StartSpanFromContext(ctx context.Context, operationName string) (*Span, context.Context) {
	var sp *Span

	otSpan, ctx := opentracing.StartSpanFromContext(ctx, operationName)
	if sp = SpanFromContext(ctx); sp == nil {
		return newTaskSpan(ctx, operationName, otSpan)
	}
	return newRegionSpan(ctx, operationName, otSpan)
}

// SpanFromContext returns the current span from the context. It returns nil if none exists.
func SpanFromContext(ctx context.Context) *Span {
	val := ctx.Value(activeSpanKey)
	if sp, ok := val.(*Span); ok {
		return sp
	}
	return nil
}

// SetTag adds a tag to the OpenTracing span and logs it as an event with the runtime tracer.
func (s *Span) SetTag(ctx context.Context, key string, value interface{}) *Span {
	if str, ok := value.(string); ok {
		log(ctx, key, str)
	}
	if s.otSpan != nil {
		s.otSpan.SetTag(key, value)
	}
	return s
}

// Finish finalizes the OpenTracing span and the current task or region.
func (s *Span) Finish() {
	for _, f := range s.finishers {
		f()
	}
}
