// Copyright 2021 The phy-go authors
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

package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/sacloud/phy-go/fake"
	"github.com/sacloud/phy-go/openapi"
	"github.com/sacloud/phy-go/pointer"
	"github.com/stretchr/testify/require"
)

var server = func() *Server {
	raidOverallStatus := openapi.RaidStatusOverallStatusOk
	return &Server{
		Engine: &fake.Engine{
			Servers: []*fake.Server{
				{
					Server: &openapi.Server{
						CachedPowerStatus: &openapi.CachedPowerStatus{
							Status: openapi.CachedPowerStatusStatusOn,
							Stored: time.Now(),
						},
						Ipv4: &openapi.ServerIpv4Global{
							GatewayAddress: "192.0.2.1",
							IpAddress:      "192.0.2.11",
							NameServers:    []string{"198.51.100.1", "198.51.100.2"},
							NetworkAddress: "192.0.2.0",
							PrefixLength:   24,
							Type:           openapi.ServerIpv4GlobalTypeCommonIpAddress,
						},
						LockStatus: nil,
						PortChannels: []openapi.PortChannel{
							{
								BondingType:   openapi.BondingTypeLacp,
								LinkSpeedType: openapi.PortChannelLinkSpeedTypeN1gbe,
								Locked:        false,
								PortChannelId: 1001,
								Ports:         []int{2001},
							},
						},
						Ports: []openapi.InterfacePort{
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
						Service: openapi.ServiceQuiet{
							Activated:   time.Now(),
							Description: nil,
							Nickname:    "server01",
							ServiceId:   "100000000001",
							Tags:        nil,
						},
						Spec: openapi.ServerSpec{
							CpuClockSpeed:         3,
							CpuCoreCount:          4,
							CpuCount:              1,
							CpuModelName:          "E3-1220 v6",
							MemorySize:            8,
							PortChannel10gbeCount: 0,
							PortChannel1gbeCount:  1,
							Storages: []openapi.Storage{
								{
									BusType:     openapi.StorageBusTypeSata,
									DeviceCount: 2,
									MediaType:   openapi.StorageMediaTypeSsd,
									Size:        1000,
								},
							},
							TotalStorageDeviceCount: 1,
						},
						Zone: openapi.Zone{
							Region: "is",
							ZoneId: 302,
						},
					},
					RaidStatus: &openapi.RaidStatus{
						LogicalVolumes: []openapi.RaidLogicalVolume{
							{
								PhysicalDeviceIds: []string{"0", "1"},
								RaidLevel:         "1",
								Status:            openapi.RaidLogicalVolumeStatusOk,
								VolumeId:          "0",
							},
						},
						Monitored:     time.Now(),
						OverallStatus: &raidOverallStatus,
						PhysicalDevices: []openapi.RaidPhysicalDevice{
							{
								DeviceId: "0",
								Slot:     0,
								Status:   openapi.RaidPhysicalDeviceStatusOk,
							},
							{
								DeviceId: "1",
								Slot:     1,
								Status:   openapi.RaidPhysicalDeviceStatusOk,
							},
						},
					},
					OSImages: []*openapi.OsImage{
						{
							ManualPartition: true,
							Name:            "Usacloud Linux",
							OsImageId:       "usacloud",
							RequirePassword: true,
							SuperuserName:   "root",
						},
					},
					PowerStatus: &openapi.ServerPowerStatus{
						Status: openapi.ServerPowerStatusStatusOn,
					},
					TrafficGraph: &openapi.TrafficGraph{
						Receive: []openapi.TrafficGraphData{
							{
								Timestamp: time.Now(),
								Value:     1,
							},
						},
						Transmit: []openapi.TrafficGraphData{
							{
								Timestamp: time.Now(),
								Value:     1,
							},
						},
					},
				},
				{
					Server: &openapi.Server{
						CachedPowerStatus: &openapi.CachedPowerStatus{
							Status: openapi.CachedPowerStatusStatusOn,
							Stored: time.Now(),
						},
						Ipv4: &openapi.ServerIpv4Global{
							GatewayAddress: "192.0.2.1",
							IpAddress:      "192.0.2.11",
							NameServers:    []string{"198.51.100.1", "198.51.100.2"},
							NetworkAddress: "192.0.2.0",
							PrefixLength:   24,
							Type:           openapi.ServerIpv4GlobalTypeCommonIpAddress,
						},
						LockStatus: nil,
						PortChannels: []openapi.PortChannel{
							{
								BondingType:   openapi.BondingTypeLacp,
								LinkSpeedType: openapi.PortChannelLinkSpeedTypeN1gbe,
								Locked:        false,
								PortChannelId: 1002,
								Ports:         []int{2002},
							},
						},
						Ports: []openapi.InterfacePort{
							{
								Enabled:             false,
								GlobalBandwidthMbps: nil,
								Internet:            nil,
								LocalBandwidthMbps:  nil,
								Mode:                nil,
								Nickname:            "server02-port01",
								PortChannelId:       1002,
								PortId:              2002,
								PrivateNetworks:     nil,
							},
						},
						ServerId: "100000000002",
						Service: openapi.ServiceQuiet{
							Activated:   time.Now(),
							Description: nil,
							Nickname:    "server02",
							ServiceId:   "100000000002",
							Tags:        nil,
						},
						Spec: openapi.ServerSpec{
							CpuClockSpeed:         3,
							CpuCoreCount:          4,
							CpuCount:              1,
							CpuModelName:          "E3-1220 v6",
							MemorySize:            8,
							PortChannel10gbeCount: 0,
							PortChannel1gbeCount:  1,
							Storages: []openapi.Storage{
								{
									BusType:     openapi.StorageBusTypeSata,
									DeviceCount: 2,
									MediaType:   openapi.StorageMediaTypeSsd,
									Size:        1000,
								},
							},
							TotalStorageDeviceCount: 1,
						},
						Zone: openapi.Zone{
							Region: "is",
							ZoneId: 302,
						},
					},
					RaidStatus: &openapi.RaidStatus{
						LogicalVolumes: []openapi.RaidLogicalVolume{
							{
								PhysicalDeviceIds: []string{"0", "1"},
								RaidLevel:         "1",
								Status:            openapi.RaidLogicalVolumeStatusOk,
								VolumeId:          "0",
							},
						},
						Monitored:     time.Now(),
						OverallStatus: &raidOverallStatus,
						PhysicalDevices: []openapi.RaidPhysicalDevice{
							{
								DeviceId: "0",
								Slot:     0,
								Status:   openapi.RaidPhysicalDeviceStatusOk,
							},
							{
								DeviceId: "1",
								Slot:     1,
								Status:   openapi.RaidPhysicalDeviceStatusOk,
							},
						},
					},
					OSImages: []*openapi.OsImage{
						{
							ManualPartition: true,
							Name:            "Usacloud Linux2",
							OsImageId:       "usacloud2",
							RequirePassword: true,
							SuperuserName:   "root",
						},
					},
					PowerStatus: &openapi.ServerPowerStatus{
						Status: openapi.ServerPowerStatusStatusOn,
					},
					TrafficGraph: &openapi.TrafficGraph{
						Receive: []openapi.TrafficGraphData{
							{
								Timestamp: time.Now(),
								Value:     1,
							},
						},
						Transmit: []openapi.TrafficGraphData{
							{
								Timestamp: time.Now(),
								Value:     1,
							},
						},
					},
				},
			},
			DedicatedSubnets: []*openapi.DedicatedSubnet{
				{
					ConfigStatus:      openapi.DedicatedSubnetConfigStatusOperational,
					DedicatedSubnetId: "100000000001",
					Firewall:          nil,
					Ipv4: openapi.Ipv4{
						BroadcastAddress:    "192.0.2.239",
						GatewayAddress:      "192.0.2.225",
						NetworkAddress:      "192.0.2.224",
						PrefixLength:        28,
						SpecialUseAddresses: nil,
					},
					LoadBalancer: nil,
					ServerCount:  1,
					Service: openapi.ServiceQuiet{
						Activated: time.Now(),
						Nickname:  "global-network01",
						ServiceId: "100000000001",
						Tags:      nil,
					},
					Zone: openapi.Zone{
						Region: "is",
						ZoneId: 302,
					},
				},
			},
			PrivateNetworks: []*openapi.PrivateNetwork{
				{
					PrivateNetworkId: "100000000001",
					ServerCount:      1,
					Service: openapi.ServiceQuiet{
						Activated: time.Now(),
						Nickname:  "private-network01",
						ServiceId: "100000000001",
					},
					VlanId: 1,
					Zone: openapi.Zone{
						Region: "is",
						ZoneId: 302,
					},
				},
			},
			Services: []*openapi.Service{
				{
					Activated:   time.Now(),
					Description: pointer.String("description1"),
					Nickname:    "nickname1",
					OptionPlans: nil,
					Plan: &openapi.ServicePlan{
						Name:   "plan-01",
						PlanId: "maker-series-spec-region-01",
					},
					ProductCategory: openapi.ServiceProductCategoryServer,
					ServiceId:       "100000000001",
					Tags: []openapi.Tag{
						{
							Color: pointer.String("ffffff"),
							Label: "label",
							TagId: 1,
						},
					},
				},
			},
		},
	}
}()

var testServer *httptest.Server

func TestMain(m *testing.M) {
	testServer = httptest.NewServer(server.Handler())
	defer testServer.Close()

	os.Exit(m.Run())
}

func TestServer_ping(t *testing.T) {
	resp, err := http.Get(testServer.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, "pong", string(body))
}

func TestServer_ListServices(t *testing.T) {
	resp, err := http.Get(testServer.URL + "/services/")
	if err != nil {
		t.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var services openapi.Services
	if err := json.Unmarshal(body, &services); err != nil {
		t.Fatal(err)
	}

	require.Equal(t, 1, services.Meta.Count)
	require.Equal(t, "nickname1", services.Services[0].Nickname)
}

func TestServer_ServerPowerControl(t *testing.T) {
	params := bytes.NewBufferString(`{"operation":"off"}`)
	req, err := http.NewRequest(http.MethodPost, testServer.URL+"/servers/100000000001/power_control/", params)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, http.StatusNoContent, resp.StatusCode)
	time.Sleep(500 * time.Millisecond) // シャットダウンされるまで少し待つ

	resp, err = http.Get(testServer.URL + "/servers/100000000001/power_status")
	if err != nil {
		t.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var status openapi.ServerPowerStatus
	if err := json.Unmarshal(body, &status); err != nil {
		t.Fatal(err)
	}

	require.Equal(t, "off", string(status.Status))
}
