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
	"reflect"

	"github.com/pkg/errors"
	config_util "github.com/prometheus/common/config"
)

var (
	// DefaultKubernetes is the default Kubernetes SD configuration
	DefaultKubernetes = Kubernetes{}
)

// KubernetesRole is role of the service in Kubernetes.
type KubernetesRole string

// The valid options for Role.
const (
	RoleNode     KubernetesRole = "node"
	RolePod      KubernetesRole = "pod"
	RoleService  KubernetesRole = "service"
	RoleEndpoint KubernetesRole = "endpoints"
	RoleIngress  KubernetesRole = "ingress"
)

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *KubernetesRole) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := unmarshal((*string)(c)); err != nil {
		return err
	}
	switch *c {
	case RoleNode, RolePod, RoleService, RoleEndpoint, RoleIngress:
		return nil
	default:
		return errors.Errorf("unknown Kubernetes SD role %q", *c)
	}
}

// Kubernetes is the configuration for Kubernetes service discovery.
type Kubernetes struct {
	APIServer          config_util.URL              `yaml:"api_server,omitempty"`
	Role               KubernetesRole               `yaml:"role"`
	HTTPClientConfig   config_util.HTTPClientConfig `yaml:",inline"`
	NamespaceDiscovery NamespaceDiscovery           `yaml:"namespaces,omitempty"`
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *Kubernetes) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = Kubernetes{}
	type plain Kubernetes
	err := unmarshal((*plain)(c))
	if err != nil {
		return err
	}
	if c.Role == "" {
		return errors.Errorf("role missing (one of: pod, service, endpoints, node, ingress)")
	}
	err = c.HTTPClientConfig.Validate()
	if err != nil {
		return err
	}
	if c.APIServer.URL == nil && !reflect.DeepEqual(c.HTTPClientConfig, config_util.HTTPClientConfig{}) {
		return errors.Errorf("to use custom HTTP client configuration please provide the 'api_server' URL explicitly")
	}
	return nil
}

// NamespaceDiscovery is the configuration for discovering
// Kubernetes namespaces.
type NamespaceDiscovery struct {
	Names []string `yaml:"names"`
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *NamespaceDiscovery) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = NamespaceDiscovery{}
	type plain NamespaceDiscovery
	return unmarshal((*plain)(c))
}
