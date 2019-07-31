// Copyright 2017 The Prometheus Authors
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
	"time"

	"github.com/pkg/errors"
	config_util "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
)

// DefaultTriton is the default Triton SD configuration.
var DefaultTriton = Triton{
	Port:            9163,
	RefreshInterval: model.Duration(60 * time.Second),
	Version:         1,
}

// Triton is the configuration for Triton based service discovery.
type Triton struct {
	Account         string                `yaml:"account"`
	DNSSuffix       string                `yaml:"dns_suffix"`
	Endpoint        string                `yaml:"endpoint"`
	Groups          []string              `yaml:"groups,omitempty"`
	Port            int                   `yaml:"port"`
	RefreshInterval model.Duration        `yaml:"refresh_interval,omitempty"`
	TLSConfig       config_util.TLSConfig `yaml:"tls_config,omitempty"`
	Version         int                   `yaml:"version"`
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *Triton) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultTriton
	type plain Triton
	err := unmarshal((*plain)(c))
	if err != nil {
		return err
	}
	if c.Account == "" {
		return errors.New("triton SD configuration requires an account")
	}
	if c.DNSSuffix == "" {
		return errors.New("triton SD configuration requires a dns_suffix")
	}
	if c.Endpoint == "" {
		return errors.New("triton SD configuration requires an endpoint")
	}
	if c.RefreshInterval <= 0 {
		return errors.New("triton SD configuration requires RefreshInterval to be a positive integer")
	}
	return nil
}
