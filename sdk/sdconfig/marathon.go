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
	"time"

	"github.com/pkg/errors"
	config_util "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
)

// DefaultMarathon is the default Marathon SD configuration.
var DefaultMarathon = Marathon{
	RefreshInterval: model.Duration(30 * time.Second),
}

// Marathon is the configuration for services running on Marathon.
type Marathon struct {
	Servers          []string                     `yaml:"servers,omitempty"`
	RefreshInterval  model.Duration               `yaml:"refresh_interval,omitempty"`
	AuthToken        config_util.Secret           `yaml:"auth_token,omitempty"`
	AuthTokenFile    string                       `yaml:"auth_token_file,omitempty"`
	HTTPClientConfig config_util.HTTPClientConfig `yaml:",inline"`
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *Marathon) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultMarathon
	type plain Marathon
	err := unmarshal((*plain)(c))
	if err != nil {
		return err
	}
	if len(c.Servers) == 0 {
		return errors.New("marathon_sd: must contain at least one Marathon server")
	}
	if len(c.AuthToken) > 0 && len(c.AuthTokenFile) > 0 {
		return errors.New("marathon_sd: at most one of auth_token & auth_token_file must be configured")
	}
	if c.HTTPClientConfig.BasicAuth != nil && (len(c.AuthToken) > 0 || len(c.AuthTokenFile) > 0) {
		return errors.New("marathon_sd: at most one of basic_auth, auth_token & auth_token_file must be configured")
	}
	if (len(c.HTTPClientConfig.BearerToken) > 0 || len(c.HTTPClientConfig.BearerTokenFile) > 0) && (len(c.AuthToken) > 0 || len(c.AuthTokenFile) > 0) {
		return errors.New("marathon_sd: at most one of bearer_token, bearer_token_file, auth_token & auth_token_file must be configured")
	}
	return c.HTTPClientConfig.Validate()
}
