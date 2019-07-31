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
	config_util "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
)

const (
	authMethodOAuth           = "OAuth"
	authMethodManagedIdentity = "ManagedIdentity"
)

// DefaultAzure is the default Azure SD configuration.
var DefaultAzure = Azure{
	Port:            80,
	RefreshInterval: model.Duration(5 * time.Minute),
	//FIXME: find a way to get this value without pulling the azure package.
	//Environment:          azure.PublicCloud.Name,
	Environment:          "AzurePublicCloud",
	AuthenticationMethod: authMethodOAuth,
}

// AzureConfig is the configuration for Azure based service discovery.
type Azure struct {
	Environment          string             `yaml:"environment,omitempty"`
	Port                 int                `yaml:"port"`
	SubscriptionID       string             `yaml:"subscription_id"`
	TenantID             string             `yaml:"tenant_id,omitempty"`
	ClientID             string             `yaml:"client_id,omitempty"`
	ClientSecret         config_util.Secret `yaml:"client_secret,omitempty"`
	RefreshInterval      model.Duration     `yaml:"refresh_interval,omitempty"`
	AuthenticationMethod string             `yaml:"authentication_method,omitempty"`
}

func validateAuthParam(param, name string) error {
	if len(param) == 0 {
		return errors.Errorf("azure SD configuration requires a %s", name)
	}
	return nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *Azure) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultAzure
	type plain Azure
	err := unmarshal((*plain)(c))
	if err != nil {
		return err
	}

	if err = validateAuthParam(c.SubscriptionID, "subscription_id"); err != nil {
		return err
	}

	if c.AuthenticationMethod == authMethodOAuth {
		if err = validateAuthParam(c.TenantID, "tenant_id"); err != nil {
			return err
		}
		if err = validateAuthParam(c.ClientID, "client_id"); err != nil {
			return err
		}
		if err = validateAuthParam(string(c.ClientSecret), "client_secret"); err != nil {
			return err
		}
	}

	if c.AuthenticationMethod != authMethodOAuth && c.AuthenticationMethod != authMethodManagedIdentity {
		return errors.Errorf("unknown authentication_type %q. Supported types are %q or %q", c.AuthenticationMethod, authMethodOAuth, authMethodManagedIdentity)
	}

	return nil
}
