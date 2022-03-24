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

package phy

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	client "github.com/sacloud/api-client-go"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/sacloud/phy-api-go/fake"
	"github.com/sacloud/phy-api-go/fake/server"
	"github.com/sacloud/phy-api-go/pointer"
)

var testServerURL = DefaultAPIRootURL

func TestMain(m *testing.M) {
	if !isAcceptanceTest() {
		sv := httptest.NewServer(testServer.Handler())
		defer sv.Close()

		testServerURL = sv.URL
	}
	os.Exit(m.Run())
}

func isAcceptanceTest() bool {
	return os.Getenv("TESTACC") != ""
}

func onlyAcceptanceTest(t *testing.T) {
	if !isAcceptanceTest() {
		t.Skip("this test will only run when 'TESTACC' is set")
	}
}

func onlyUnitTest(t *testing.T) {
	if isAcceptanceTest() {
		t.Skip("this test will only run when 'TESTACC' is not set")
	}
}

func skipIfEnvIsSet(t *testing.T, envKey string) {
	if os.Getenv(envKey) != "" {
		t.Skipf("environment variable %q is set, skip.", envKey)
	}
}

var testHTTPClient = &http.Client{}

func testClient(t *testing.T) *Client {
	token := "dummy"
	secret := "dummy"
	if isAcceptanceTest() {
		token = os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
		secret = os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

		if token == "" || secret == "" {
			t.Fatal("environment variable must be set: SAKURACLOUD_ACCESS_TOKEN, SAKURACLOUD_ACCESS_TOKEN_SECRET")
		}
	}

	return &Client{
		APIRootURL: testServerURL,
		Options: &client.Options{
			AccessToken:       token,
			AccessTokenSecret: secret,
			HttpClient:        testHTTPClient,
		},
	}
}

var raidOverallStatus = v1.RaidStatusOverallStatusOk

var testServer = &server.Server{
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
					ServerId: "400000000001",
					Service: v1.ServiceQuiet{
						Activated:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						Description: nil,
						Nickname:    "server01",
						ServiceId:   "400000000001",
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
				RaidStatus: &v1.RaidStatus{
					LogicalVolumes: []v1.RaidLogicalVolume{
						{
							PhysicalDeviceIds: []string{"0", "1"},
							RaidLevel:         "1",
							Status:            v1.RaidLogicalVolumeStatusOk,
							VolumeId:          "0",
						},
					},
					Monitored:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					OverallStatus: &raidOverallStatus,
					PhysicalDevices: []v1.RaidPhysicalDevice{
						{
							DeviceId: "0",
							Slot:     0,
							Status:   v1.RaidPhysicalDeviceStatusOk,
						},
						{
							DeviceId: "1",
							Slot:     1,
							Status:   v1.RaidPhysicalDeviceStatusOk,
						},
					},
				},
				OSImages: []*v1.OsImage{
					{
						ManualPartition: true,
						Name:            "Usacloud Linux",
						OsImageId:       "usacloud",
						RequirePassword: true,
						SuperuserName:   "root",
					},
				},
				PowerStatus: &v1.ServerPowerStatus{
					Status: v1.ServerPowerStatusStatusOn,
				},
				TrafficGraph: &v1.TrafficGraph{
					Receive: []v1.TrafficGraphData{
						{
							Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
							Value:     1,
						},
					},
					Transmit: []v1.TrafficGraphData{
						{
							Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
							Value:     1,
						},
					},
				},
			},
		},
		Services: []*v1.Service{
			{
				Activated:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				Description: pointer.String("description01"),
				Nickname:    "service01",
				OptionPlans: nil,
				Plan: &v1.ServicePlan{
					Name:   "plan01",
					PlanId: "maker-series-spec-region-01",
				},
				ProductCategory: v1.ServiceProductCategoryServer,
				ServiceId:       "100000000001",
				Tags: []v1.Tag{
					{
						Color: pointer.String("tag1"),
						Label: "label",
						TagId: 1,
					},
				},
			},
		},
		DedicatedSubnets: []*v1.DedicatedSubnet{
			{
				ConfigStatus:      v1.DedicatedSubnetConfigStatusOperational,
				DedicatedSubnetId: "2000000000001",
				Firewall:          nil,
				Ipv4: v1.Ipv4{
					BroadcastAddress: "192.0.2.239",
					GatewayAddress:   "192.0.2.225",
					NetworkAddress:   "192.0.2.224",
					PrefixLength:     28,
				},
				Ipv6: v1.Ipv6{
					Enabled: false,
				},
				LoadBalancer: nil,
				ServerCount:  1,
				Service: v1.ServiceQuiet{
					Activated:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Description: pointer.String("description01"),
					Nickname:    "dedicated_subnet01",
					ServiceId:   "200000000001",
					Tags: &[]v1.Tag{
						{
							Color: pointer.String("tag1"),
							Label: "label",
							TagId: 1,
						},
					},
				},
				Zone: v1.Zone{
					Region: "石狩",
					ZoneId: 301,
				},
			},
		},
		PrivateNetworks: []*v1.PrivateNetwork{
			{
				PrivateNetworkId: "300000000001",
				ServerCount:      1,
				Service: v1.ServiceQuiet{
					Activated: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Nickname:  "private-network01",
					ServiceId: "300000000001",
				},
				VlanId: 1,
				Zone: v1.Zone{
					Region: "is",
					ZoneId: 302,
				},
			},
		},
	},
}
