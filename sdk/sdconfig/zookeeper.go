// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sdconfig

import (
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
)

var (
	// DefaultServerset is the default Serverset SD configuration.
	DefaultServerset = Serverset{
		Timeout: model.Duration(10 * time.Second),
	}
	// DefaultNerve is the default Nerve SD configuration.
	DefaultNerve = Nerve{
		Timeout: model.Duration(10 * time.Second),
	}
)

// Serverset is the configuration for Twitter serversets in Zookeeper based discovery.
type Serverset struct {
	Servers []string       `yaml:"servers"`
	Paths   []string       `yaml:"paths"`
	Timeout model.Duration `yaml:"timeout,omitempty"`
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *Serverset) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultServerset
	type plain Serverset
	err := unmarshal((*plain)(c))
	if err != nil {
		return err
	}
	if len(c.Servers) == 0 {
		return errors.New("serverset SD config must contain at least one Zookeeper server")
	}
	if len(c.Paths) == 0 {
		return errors.New("serverset SD config must contain at least one path")
	}
	for _, path := range c.Paths {
		if !strings.HasPrefix(path, "/") {
			return errors.Errorf("serverset SD config paths must begin with '/': %s", path)
		}
	}
	return nil
}

// Nerve is the configuration for AirBnB's Nerve in Zookeeper based discovery.
type Nerve struct {
	Servers []string       `yaml:"servers"`
	Paths   []string       `yaml:"paths"`
	Timeout model.Duration `yaml:"timeout,omitempty"`
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *Nerve) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultNerve
	type plain Nerve
	err := unmarshal((*plain)(c))
	if err != nil {
		return err
	}
	if len(c.Servers) == 0 {
		return errors.New("nerve SD config must contain at least one Zookeeper server")
	}
	if len(c.Paths) == 0 {
		return errors.New("nerve SD config must contain at least one path")
	}
	for _, path := range c.Paths {
		if !strings.HasPrefix(path, "/") {
			return errors.Errorf("nerve SD config paths must begin with '/': %s", path)
		}
	}
	return nil
}
