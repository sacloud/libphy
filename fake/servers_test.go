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

package fake

import (
	"testing"
	"time"

	"github.com/sacloud/phy-go/openapi"
	"github.com/stretchr/testify/require"
)

func TestDataStore_Servers(t *testing.T) {
	raidOverallStatus := openapi.RaidStatusOverallStatusOk
	ds := &Engine{
		Servers: []*Server{
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
	}

	t.Run("select all", func(t *testing.T) {
		servers, err := ds.ListServers(openapi.ListServersParams{})
		require.NoError(t, err)
		require.Equal(t, len(ds.Servers), servers.Meta.Count)
		require.Len(t, servers.Servers, len(ds.Servers))
	})

	t.Run("select by id", func(t *testing.T) {
		server, err := ds.ReadServer("100000000002")
		require.NoError(t, err)
		require.Equal(t, "server02", server.Service.Nickname)
	})

	t.Run("list os images", func(t *testing.T) {
		images, err := ds.ListOSImages("100000000002")
		require.NoError(t, err)
		require.Equal(t, "usacloud2", images[0].OsImageId)
	})

	t.Run("os install", func(t *testing.T) {
		err := ds.OSInstall("100000000002", openapi.OsInstallParameter{
			ManualPartition: true,
			OsImageId:       "usacloud2",
			Password:        "passw0rd",
		})
		require.NoError(t, err)
	})

	t.Run("select port-channel by id", func(t *testing.T) {
		pc, err := ds.ReadServerPortChannel("100000000001", 1001)
		require.NoError(t, err)
		require.Equal(t, 1001, pc.PortChannelId)
	})

	t.Run("select port by id", func(t *testing.T) {
		port, err := ds.ReadServerPort("100000000002", 2002)
		require.NoError(t, err)
		require.Equal(t, "server02-port01", port.Nickname)
	})

	t.Run("update server port", func(t *testing.T) {
		_, err := ds.UpdateServerPort("100000000002", 2002, openapi.UpdateServerPortParameter{
			Nickname: "server02-port01-upd",
		})
		require.NoError(t, err)

		port, err := ds.ReadServerPort("100000000002", 2002)
		require.NoError(t, err)
		require.Equal(t, "server02-port01-upd", port.Nickname)
	})

	t.Run("Enable server port", func(t *testing.T) {
		before, err := ds.ReadServerPort("100000000002", 2002)
		require.NoError(t, err)
		require.False(t, before.Enabled)

		_, err = ds.EnableServerPort("100000000002", 2002, openapi.EnableServerPortParameter{Enable: true})
		require.NoError(t, err)

		after, err := ds.ReadServerPort("100000000002", 2002)
		require.NoError(t, err)
		require.True(t, after.Enabled)
	})

	t.Run("assign network", func(t *testing.T) {
		t.Run("to common global", func(t *testing.T) {
			internetType := openapi.AssignNetworkParameterInternetTypeCommonSubnet
			_, err := ds.ServerAssignNetwork("100000000002", 2002, openapi.AssignNetworkParameter{
				InternetType: &internetType,
				Mode:         openapi.AssignNetworkParameterModeAccess,
			})
			require.NoError(t, err)

			port, err := ds.ReadServerPort("100000000002", 2002)
			require.NoError(t, err)
			require.Equal(t, *port.Mode, openapi.InterfacePortModeAccess)
			require.Equal(t, port.Internet.SubnetType, openapi.InternetSubnetTypeCommonSubnet)
		})

		t.Run("to dedicated subnet", func(t *testing.T) {
			internetType := openapi.AssignNetworkParameterInternetTypeDedicatedSubnet
			subnetId := "100000000001"
			_, err := ds.ServerAssignNetwork("100000000002", 2002, openapi.AssignNetworkParameter{
				InternetType:      &internetType,
				Mode:              openapi.AssignNetworkParameterModeAccess,
				DedicatedSubnetId: &subnetId,
			})
			require.NoError(t, err)

			port, err := ds.ReadServerPort("100000000002", 2002)
			require.NoError(t, err)
			require.Equal(t, *port.Mode, openapi.InterfacePortModeAccess)
			require.Equal(t, port.Internet.SubnetType, openapi.InternetSubnetTypeDedicatedSubnet)
		})

		t.Run("to private subnet", func(t *testing.T) {
			privateNetworkIds := []string{"100000000001"}
			_, err := ds.ServerAssignNetwork("100000000002", 2002, openapi.AssignNetworkParameter{
				Mode:              openapi.AssignNetworkParameterModeAccess,
				PrivateNetworkIds: &privateNetworkIds,
			})
			require.NoError(t, err)

			port, err := ds.ReadServerPort("100000000002", 2002)
			require.NoError(t, err)
			require.Equal(t, *port.Mode, openapi.InterfacePortModeAccess)
			require.Nil(t, port.Internet)
			require.NotEmpty(t, port.PrivateNetworks)
		})

		t.Run("trunk port", func(t *testing.T) {
			internetType := openapi.AssignNetworkParameterInternetTypeCommonSubnet
			privateNetworkIds := []string{"100000000001"}
			_, err := ds.ServerAssignNetwork("100000000002", 2002, openapi.AssignNetworkParameter{
				InternetType:      &internetType,
				Mode:              openapi.AssignNetworkParameterModeTrunk,
				PrivateNetworkIds: &privateNetworkIds,
			})
			require.NoError(t, err)

			port, err := ds.ReadServerPort("100000000002", 2002)
			require.NoError(t, err)
			require.Equal(t, *port.Mode, openapi.InterfacePortModeTrunk)
			require.Equal(t, port.Internet.SubnetType, openapi.InternetSubnetTypeCommonSubnet)
			require.NotEmpty(t, port.PrivateNetworks)
		})

		t.Run("disconnect", func(t *testing.T) {
			_, err := ds.ServerAssignNetwork("100000000002", 2002, openapi.AssignNetworkParameter{
				Mode: openapi.AssignNetworkParameterModeAccess,
			})
			require.NoError(t, err)

			port, err := ds.ReadServerPort("100000000002", 2002)
			require.NoError(t, err)
			require.Equal(t, *port.Mode, openapi.InterfacePortModeAccess)
			require.Nil(t, port.Internet)
			require.Empty(t, port.PrivateNetworks)
		})
	})

	t.Run("power control", func(t *testing.T) {
		err := ds.ServerPowerControl("100000000002", openapi.PowerControlParameter{Operation: "reset"})
		require.NoError(t, err)
		time.Sleep(ds.actionInterval() * 10) // 反映されるまで少し待つ

		powerStatus, err := ds.ReadServerPowerStatus("100000000002")
		require.NoError(t, err)
		require.Equal(t, openapi.ServerPowerStatusStatusOn, powerStatus.Status)

		err = ds.ServerPowerControl("100000000002", openapi.PowerControlParameter{Operation: "off"})
		require.NoError(t, err)
		time.Sleep(ds.actionInterval() * 10) // 反映されるまで少し待つ

		powerStatus, err = ds.ReadServerPowerStatus("100000000002")
		require.NoError(t, err)
		require.Equal(t, openapi.ServerPowerStatusStatusOff, powerStatus.Status)
	})

	t.Run("read RAID status", func(t *testing.T) {
		status, err := ds.ReadRAIDStatus("100000000002", openapi.ReadRAIDStatusParams{Refresh: nil})
		require.NoError(t, err)
		require.NotNil(t, status.OverallStatus)
		require.Equal(t, *status.OverallStatus, openapi.RaidStatusOverallStatusOk)
	})

	// 副作用(portIdが変わる)があるため最後に実施
	t.Run("configure bonding", func(t *testing.T) {
		t.Run("LACP", func(t *testing.T) {
			pc, err := ds.ServerConfigureBonding("100000000002", 1002, openapi.ConfigureBondingParameter{
				BondingType:   openapi.BondingTypeLacp,
				PortNicknames: nil,
			})
			require.NoError(t, err)
			require.Len(t, pc.Ports, 1)

			server := ds.getServerById("100000000002")
			require.Len(t, server.Server.Ports, 1)
		})
		t.Run("Single", func(t *testing.T) {
			pc, err := ds.ServerConfigureBonding("100000000002", 1002, openapi.ConfigureBondingParameter{
				BondingType:   openapi.BondingTypeSingle,
				PortNicknames: nil,
			})
			require.NoError(t, err)
			require.Len(t, pc.Ports, 2)

			server := ds.getServerById("100000000002")
			require.Len(t, server.Server.Ports, 2)
		})
	})
}
