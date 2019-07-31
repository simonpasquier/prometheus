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

// DefaultOpenStack is the default OpenStack SD configuration.
var DefaultOpenStack = OpenStack{
	Port:            80,
	RefreshInterval: model.Duration(60 * time.Second),
}

// OpenStack is the configuration for OpenStack based service discovery.
type OpenStack struct {
	IdentityEndpoint            string                `yaml:"identity_endpoint"`
	Username                    string                `yaml:"username"`
	UserID                      string                `yaml:"userid"`
	Password                    config_util.Secret    `yaml:"password"`
	ProjectName                 string                `yaml:"project_name"`
	ProjectID                   string                `yaml:"project_id"`
	DomainName                  string                `yaml:"domain_name"`
	DomainID                    string                `yaml:"domain_id"`
	ApplicationCredentialName   string                `yaml:"application_credential_name"`
	ApplicationCredentialID     string                `yaml:"application_credential_id"`
	ApplicationCredentialSecret config_util.Secret    `yaml:"application_credential_secret"`
	Role                        OpenStackRole         `yaml:"role"`
	Region                      string                `yaml:"region"`
	RefreshInterval             model.Duration        `yaml:"refresh_interval,omitempty"`
	Port                        int                   `yaml:"port"`
	AllTenants                  bool                  `yaml:"all_tenants,omitempty"`
	TLSConfig                   config_util.TLSConfig `yaml:"tls_config,omitempty"`
}

// OpenStackRole is the role of the target in OpenStack.
type OpenStackRole string

// The valid options for OpenStackRole.
const (
	// OpenStack document reference
	// https://docs.openstack.org/nova/pike/admin/arch.html#hypervisors
	OpenStackRoleHypervisor OpenStackRole = "hypervisor"
	// OpenStack document reference
	// https://docs.openstack.org/horizon/pike/user/launch-instances.html
	OpenStackRoleInstance OpenStackRole = "instance"
)

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *OpenStackRole) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := unmarshal((*string)(c)); err != nil {
		return err
	}
	switch *c {
	case OpenStackRoleHypervisor, OpenStackRoleInstance:
		return nil
	default:
		return errors.Errorf("unknown OpenStack SD role %q", *c)
	}
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *OpenStack) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultOpenStack
	type plain OpenStack
	err := unmarshal((*plain)(c))
	if err != nil {
		return err
	}
	if c.Role == "" {
		return errors.New("role missing (one of: instance, hypervisor)")
	}
	if c.Region == "" {
		return errors.New("openstack SD configuration requires a region")
	}
	return nil
}
