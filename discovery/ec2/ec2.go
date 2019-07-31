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

package ec2

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	"github.com/prometheus/common/model"

	"github.com/simonpasquier/prometheus/discovery/refresh"
	"github.com/simonpasquier/prometheus/sdk/sdconfig"
	"github.com/simonpasquier/prometheus/sdk/strutil"
	"github.com/simonpasquier/prometheus/sdk/targetgroup"
)

const (
	ec2Label                = model.MetaLabelPrefix + "ec2_"
	ec2LabelAZ              = ec2Label + "availability_zone"
	ec2LabelInstanceID      = ec2Label + "instance_id"
	ec2LabelInstanceState   = ec2Label + "instance_state"
	ec2LabelInstanceType    = ec2Label + "instance_type"
	ec2LabelOwnerID         = ec2Label + "owner_id"
	ec2LabelPlatform        = ec2Label + "platform"
	ec2LabelPublicDNS       = ec2Label + "public_dns_name"
	ec2LabelPublicIP        = ec2Label + "public_ip"
	ec2LabelPrivateDNS      = ec2Label + "private_dns_name"
	ec2LabelPrivateIP       = ec2Label + "private_ip"
	ec2LabelPrimarySubnetID = ec2Label + "primary_subnet_id"
	ec2LabelSubnetID        = ec2Label + "subnet_id"
	ec2LabelTag             = ec2Label + "tag_"
	ec2LabelVPCID           = ec2Label + "vpc_id"
	subnetSeparator         = ","
)

// Discovery periodically performs EC2-SD requests. It implements
// the Discoverer interface.
type Discovery struct {
	*refresh.Discovery
	aws      *aws.Config
	interval time.Duration
	profile  string
	roleARN  string
	port     int
	filters  []*sdconfig.Filter
}

// NewDiscovery returns a new EC2Discovery which periodically refreshes its targets.
func NewDiscovery(conf *sdconfig.EC2, logger log.Logger) *Discovery {
	creds := credentials.NewStaticCredentials(conf.AccessKey, string(conf.SecretKey), "")
	if conf.AccessKey == "" && conf.SecretKey == "" {
		creds = nil
	}
	if logger == nil {
		logger = log.NewNopLogger()
	}
	d := &Discovery{
		aws: &aws.Config{
			Endpoint:    &conf.Endpoint,
			Region:      &conf.Region,
			Credentials: creds,
		},
		profile:  conf.Profile,
		roleARN:  conf.RoleARN,
		filters:  conf.Filters,
		interval: time.Duration(conf.RefreshInterval),
		port:     conf.Port,
	}
	d.Discovery = refresh.NewDiscovery(
		logger,
		"ec2",
		time.Duration(conf.RefreshInterval),
		d.refresh,
	)
	return d
}

func (d *Discovery) refresh(ctx context.Context) ([]*targetgroup.Group, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  *d.aws,
		Profile: d.profile,
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not create aws session")
	}

	var ec2s *ec2.EC2
	if d.roleARN != "" {
		creds := stscreds.NewCredentials(sess, d.roleARN)
		ec2s = ec2.New(sess, &aws.Config{Credentials: creds})
	} else {
		ec2s = ec2.New(sess)
	}
	tg := &targetgroup.Group{
		Source: *d.aws.Region,
	}

	var filters []*ec2.Filter
	for _, f := range d.filters {
		filters = append(filters, &ec2.Filter{
			Name:   aws.String(f.Name),
			Values: aws.StringSlice(f.Values),
		})
	}

	input := &ec2.DescribeInstancesInput{Filters: filters}

	if err = ec2s.DescribeInstancesPagesWithContext(ctx, input, func(p *ec2.DescribeInstancesOutput, lastPage bool) bool {
		for _, r := range p.Reservations {
			for _, inst := range r.Instances {
				if inst.PrivateIpAddress == nil {
					continue
				}
				labels := model.LabelSet{
					ec2LabelInstanceID: model.LabelValue(*inst.InstanceId),
				}

				if r.OwnerId != nil {
					labels[ec2LabelOwnerID] = model.LabelValue(*r.OwnerId)
				}

				labels[ec2LabelPrivateIP] = model.LabelValue(*inst.PrivateIpAddress)
				if inst.PrivateDnsName != nil {
					labels[ec2LabelPrivateDNS] = model.LabelValue(*inst.PrivateDnsName)
				}
				addr := net.JoinHostPort(*inst.PrivateIpAddress, fmt.Sprintf("%d", d.port))
				labels[model.AddressLabel] = model.LabelValue(addr)

				if inst.Platform != nil {
					labels[ec2LabelPlatform] = model.LabelValue(*inst.Platform)
				}

				if inst.PublicIpAddress != nil {
					labels[ec2LabelPublicIP] = model.LabelValue(*inst.PublicIpAddress)
					labels[ec2LabelPublicDNS] = model.LabelValue(*inst.PublicDnsName)
				}

				labels[ec2LabelAZ] = model.LabelValue(*inst.Placement.AvailabilityZone)
				labels[ec2LabelInstanceState] = model.LabelValue(*inst.State.Name)
				labels[ec2LabelInstanceType] = model.LabelValue(*inst.InstanceType)

				if inst.VpcId != nil {
					labels[ec2LabelVPCID] = model.LabelValue(*inst.VpcId)
					labels[ec2LabelPrimarySubnetID] = model.LabelValue(*inst.SubnetId)

					// Deduplicate VPC Subnet IDs maintaining the order of the network interfaces returned by EC2.
					var subnets []string
					subnetsMap := make(map[string]struct{})
					for _, eni := range inst.NetworkInterfaces {
						if eni.SubnetId == nil {
							continue
						}
						if _, ok := subnetsMap[*eni.SubnetId]; !ok {
							subnetsMap[*eni.SubnetId] = struct{}{}
							subnets = append(subnets, *eni.SubnetId)
						}
					}
					labels[ec2LabelSubnetID] = model.LabelValue(
						subnetSeparator +
							strings.Join(subnets, subnetSeparator) +
							subnetSeparator)
				}

				for _, t := range inst.Tags {
					if t == nil || t.Key == nil || t.Value == nil {
						continue
					}
					name := strutil.SanitizeLabelName(*t.Key)
					labels[ec2LabelTag+model.LabelName(name)] = model.LabelValue(*t.Value)
				}
				tg.Targets = append(tg.Targets, labels)
			}
		}
		return true
	}); err != nil {
		return nil, errors.Wrap(err, "could not describe instances")
	}
	return []*targetgroup.Group{tg}, nil
}
