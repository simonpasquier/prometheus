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

package openstack

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	conntrack "github.com/mwitkow/go-conntrack"
	"github.com/pkg/errors"
	config_util "github.com/prometheus/common/config"

	"github.com/simonpasquier/prometheus/discovery/refresh"
	"github.com/simonpasquier/prometheus/sdk/sdconfig"
	"github.com/simonpasquier/prometheus/sdk/targetgroup"
)

type refresher interface {
	refresh(context.Context) ([]*targetgroup.Group, error)
}

// NewDiscovery returns a new OpenStack Discoverer which periodically refreshes its targets.
func NewDiscovery(conf *sdconfig.OpenStack, l log.Logger) (*refresh.Discovery, error) {
	r, err := newRefresher(conf, l)
	if err != nil {
		return nil, err
	}
	return refresh.NewDiscovery(
		l,
		"openstack",
		time.Duration(conf.RefreshInterval),
		r.refresh,
	), nil

}

func newRefresher(conf *sdconfig.OpenStack, l log.Logger) (refresher, error) {
	var opts gophercloud.AuthOptions
	if conf.IdentityEndpoint == "" {
		var err error
		opts, err = openstack.AuthOptionsFromEnv()
		if err != nil {
			return nil, err
		}
	} else {
		opts = gophercloud.AuthOptions{
			IdentityEndpoint:            conf.IdentityEndpoint,
			Username:                    conf.Username,
			UserID:                      conf.UserID,
			Password:                    string(conf.Password),
			TenantName:                  conf.ProjectName,
			TenantID:                    conf.ProjectID,
			DomainName:                  conf.DomainName,
			DomainID:                    conf.DomainID,
			ApplicationCredentialID:     conf.ApplicationCredentialID,
			ApplicationCredentialName:   conf.ApplicationCredentialName,
			ApplicationCredentialSecret: string(conf.ApplicationCredentialSecret),
		}
	}
	client, err := openstack.NewClient(opts.IdentityEndpoint)
	if err != nil {
		return nil, err
	}
	tls, err := config_util.NewTLSConfig(&conf.TLSConfig)
	if err != nil {
		return nil, err
	}
	client.HTTPClient = http.Client{
		Transport: &http.Transport{
			IdleConnTimeout: 5 * time.Duration(conf.RefreshInterval),
			TLSClientConfig: tls,
			DialContext: conntrack.NewDialContextFunc(
				conntrack.DialWithTracing(),
				conntrack.DialWithName("openstack_sd"),
			),
		},
		Timeout: 5 * time.Duration(conf.RefreshInterval),
	}
	switch conf.Role {
	case sdconfig.OpenStackRoleHypervisor:
		return newHypervisorDiscovery(client, &opts, conf.Port, conf.Region, l), nil
	case sdconfig.OpenStackRoleInstance:
		return newInstanceDiscovery(client, &opts, conf.Port, conf.Region, conf.AllTenants, l), nil
	}
	return nil, errors.New("unknown OpenStack discovery role")
}
