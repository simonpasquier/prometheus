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

	// FIXME: get rid of dependency on github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
	config_util "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
)

// DefaultEC2 is the default EC2 SD configuration.
var DefaultEC2 = EC2{
	Port:            80,
	RefreshInterval: model.Duration(60 * time.Second),
}

// Filter is the configuration for filtering EC2 instances.
type Filter struct {
	Name   string   `yaml:"name"`
	Values []string `yaml:"values"`
}

// EC2 is the configuration for EC2 based service discovery.
type EC2 struct {
	Endpoint        string             `yaml:"endpoint"`
	Region          string             `yaml:"region"`
	AccessKey       string             `yaml:"access_key,omitempty"`
	SecretKey       config_util.Secret `yaml:"secret_key,omitempty"`
	Profile         string             `yaml:"profile,omitempty"`
	RoleARN         string             `yaml:"role_arn,omitempty"`
	RefreshInterval model.Duration     `yaml:"refresh_interval,omitempty"`
	Port            int                `yaml:"port"`
	Filters         []*Filter          `yaml:"filters"`
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *EC2) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultEC2
	type plain EC2
	err := unmarshal((*plain)(c))
	if err != nil {
		return err
	}
	if c.Region == "" {
		sess, err := session.NewSession()
		if err != nil {
			return err
		}
		metadata := ec2metadata.New(sess)
		region, err := metadata.Region()
		if err != nil {
			return errors.New("EC2 SD configuration requires a region")
		}
		c.Region = region
	}
	for _, f := range c.Filters {
		if len(f.Values) == 0 {
			return errors.New("EC2 SD configuration filter values cannot be empty")
		}
	}
	return nil
}
