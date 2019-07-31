// Copyright 2016 The Prometheus Authors
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
	// DefaultDNS is the default DNS SD configuration.
	DefaultDNS = DNS{
		RefreshInterval: model.Duration(30 * time.Second),
		Type:            "SRV",
	}
)

// DNS is the configuration for DNS based service discovery.
type DNS struct {
	Names           []string       `yaml:"names"`
	RefreshInterval model.Duration `yaml:"refresh_interval,omitempty"`
	Type            string         `yaml:"type"`
	Port            int            `yaml:"port"` // Ignored for SRV records
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *DNS) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultDNS
	type plain DNS
	err := unmarshal((*plain)(c))
	if err != nil {
		return err
	}
	if len(c.Names) == 0 {
		return errors.New("DNS-SD config must contain at least one SRV record name")
	}
	switch strings.ToUpper(c.Type) {
	case "SRV":
	case "A", "AAAA":
		if c.Port == 0 {
			return errors.New("a port is required in DNS-SD configs for all record types except SRV")
		}
	default:
		return errors.Errorf("invalid DNS-SD records type %s", c.Type)
	}
	return nil
}
