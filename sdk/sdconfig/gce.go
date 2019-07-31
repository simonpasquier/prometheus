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
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
)

// DefaultGCE is the default GCE SD configuration.
var DefaultGCE = GCE{
	Port:            80,
	TagSeparator:    ",",
	RefreshInterval: model.Duration(60 * time.Second),
}

// GCE is the configuration for GCE based service discovery.
type GCE struct {
	// Project: The Google Cloud Project ID
	Project string `yaml:"project"`

	// Zone: The zone of the scrape targets.
	// If you need to configure multiple zones use multiple gce_sd_configs
	Zone string `yaml:"zone"`

	// Filter: Can be used optionally to filter the instance list by other criteria.
	// Syntax of this filter string is described here in the filter query parameter section:
	// https://cloud.google.com/compute/docs/reference/latest/instances/list
	Filter string `yaml:"filter,omitempty"`

	RefreshInterval model.Duration `yaml:"refresh_interval,omitempty"`
	Port            int            `yaml:"port"`
	TagSeparator    string         `yaml:"tag_separator,omitempty"`
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *GCE) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultGCE
	type plain GCE
	err := unmarshal((*plain)(c))
	if err != nil {
		return err
	}
	if c.Project == "" {
		return errors.New("GCE SD configuration requires a project")
	}
	if c.Zone == "" {
		return errors.New("GCE SD configuration requires a zone")
	}
	return nil
}
