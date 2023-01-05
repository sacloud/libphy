// Copyright 2021-2023 The phy-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fake

import (
	"testing"
	"time"

	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/stretchr/testify/require"
)

func TestDataStore_DedicatedSubnets(t *testing.T) {
	ds := &Engine{
		DedicatedSubnets: []*v1.DedicatedSubnet{
			{
				ConfigStatus:      v1.DedicatedSubnetConfigStatusOperational,
				DedicatedSubnetId: "100000000001",
				Firewall:          nil,
				Ipv4: v1.Ipv4{
					BroadcastAddress:    "192.0.2.239",
					GatewayAddress:      "192.0.2.225",
					NetworkAddress:      "192.0.2.224",
					PrefixLength:        28,
					SpecialUseAddresses: nil,
				},
				LoadBalancer: nil,
				ServerCount:  1,
				Service: v1.ServiceQuiet{
					Activated: time.Now(),
					Nickname:  "global-network01",
					ServiceId: "100000000001",
					Tags:      nil,
				},
				Zone: v1.Zone{
					Region: "is",
					ZoneId: 302,
				},
			},
			{
				ConfigStatus:      v1.DedicatedSubnetConfigStatusOperational,
				DedicatedSubnetId: "100000000002",
				Firewall:          nil,
				Ipv4: v1.Ipv4{
					BroadcastAddress:    "192.0.2.239",
					GatewayAddress:      "192.0.2.225",
					NetworkAddress:      "192.0.2.224",
					PrefixLength:        28,
					SpecialUseAddresses: nil,
				},
				LoadBalancer: nil,
				ServerCount:  1,
				Service: v1.ServiceQuiet{
					Activated: time.Now(),
					Nickname:  "global-network02",
					ServiceId: "100000000002",
					Tags:      nil,
				},
				Zone: v1.Zone{
					Region: "is",
					ZoneId: 302,
				},
			},
		},
	}

	t.Run("select all", func(t *testing.T) {
		subnets, err := ds.ListDedicatedSubnets(v1.ListDedicatedSubnetsParams{})
		require.NoError(t, err)
		require.Equal(t, len(ds.DedicatedSubnets), subnets.Meta.Count)
		require.Len(t, subnets.DedicatedSubnets, len(ds.DedicatedSubnets))
	})

	t.Run("select by id", func(t *testing.T) {
		subnet, err := ds.ReadDedicatedSubnet("100000000002", v1.ReadDedicatedSubnetParams{})
		require.NoError(t, err)
		require.Equal(t, "global-network02", subnet.Service.Nickname)
	})
}
