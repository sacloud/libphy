// Copyright 2021-2022 The phy-api-go authors
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

package phy_test

import (
	"net/http/httptest"
	"time"

	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/sacloud/phy-api-go/fake"
	"github.com/sacloud/phy-api-go/fake/server"
)

func init() {
	initFakeServer()
}

func initFakeServer() {
	sv := httptest.NewServer(fakeServer.Handler())
	serverURL = sv.URL
}

var fakeServer = &server.Server{
	Engine: &fake.Engine{
		Servers: []*fake.Server{
			{
				Server: &v1.Server{
					CachedPowerStatus: &v1.CachedPowerStatus{
						Status: v1.CachedPowerStatusStatusOn,
						Stored: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					},
					Ipv4: &v1.ServerIpv4Global{
						GatewayAddress: "192.0.2.1",
						IpAddress:      "192.0.2.11",
						NameServers:    []string{"198.51.100.1", "198.51.100.2"},
						NetworkAddress: "192.0.2.0",
						PrefixLength:   24,
						Type:           v1.ServerIpv4GlobalTypeCommonIpAddress,
					},
					LockStatus: nil,
					PortChannels: []v1.PortChannel{
						{
							BondingType:   v1.BondingTypeLacp,
							LinkSpeedType: v1.PortChannelLinkSpeedTypeN1gbe,
							Locked:        false,
							PortChannelId: 1001,
							Ports:         []int{2001},
						},
					},
					Ports: []v1.InterfacePort{
						{
							Enabled:             true,
							GlobalBandwidthMbps: nil,
							Internet:            nil,
							LocalBandwidthMbps:  nil,
							Mode:                nil,
							Nickname:            "server01-port01",
							PortChannelId:       1001,
							PortId:              2001,
							PrivateNetworks:     nil,
						},
					},
					ServerId: "100000000001",
					Service: v1.ServiceQuiet{
						Activated:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						Description: nil,
						Nickname:    "server01",
						ServiceId:   "100000000001",
						Tags:        nil,
					},
					Spec: v1.ServerSpec{
						CpuClockSpeed:         3,
						CpuCoreCount:          4,
						CpuCount:              1,
						CpuModelName:          "E3-1220 v6",
						MemorySize:            8,
						PortChannel10gbeCount: 0,
						PortChannel1gbeCount:  1,
						Storages: []v1.Storage{
							{
								BusType:     v1.StorageBusTypeSata,
								DeviceCount: 2,
								MediaType:   v1.StorageMediaTypeSsd,
								Size:        1000,
							},
						},
						TotalStorageDeviceCount: 1,
					},
					Zone: v1.Zone{
						Region: "is",
						ZoneId: 302,
					},
				},
			},
		},
		Services: []*v1.Service{
			{
				Nickname:  "server01",
				ServiceId: "100000000001",
			},
		},
	},
}
