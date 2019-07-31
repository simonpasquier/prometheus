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
	"regexp"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
)

var (
	patFileSDName = regexp.MustCompile(`^[^*]*(\*[^/]*)?\.(json|yml|yaml|JSON|YML|YAML)$`)

	// DefaultFile is the default file SD configuration.
	DefaultFile = File{
		RefreshInterval: model.Duration(5 * time.Minute),
	}
)

// File is the configuration for file based discovery.
type File struct {
	Files           []string       `yaml:"files"`
	RefreshInterval model.Duration `yaml:"refresh_interval,omitempty"`
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *File) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultFile
	type plain File
	err := unmarshal((*plain)(c))
	if err != nil {
		return err
	}
	if len(c.Files) == 0 {
		return errors.New("file service discovery config must contain at least one path name")
	}
	for _, name := range c.Files {
		if !patFileSDName.MatchString(name) {
			return errors.Errorf("path name %q is not valid for file discovery", name)
		}
	}
	return nil
}
