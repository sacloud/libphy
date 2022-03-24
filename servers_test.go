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
	"context"
	"testing"

	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/sacloud/phy-api-go/pointer"
	"github.com/stretchr/testify/require"
)

func TestServerOp_List(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()

	tests := []struct {
		name    string
		params  *v1.ListServersParams
		want    *v1.Servers
		wantErr bool
	}{
		{
			name:   "minimum",
			params: &v1.ListServersParams{},
			want: &v1.Servers{
				Meta: v1.PaginateMeta{
					Count: 1,
				},
				Servers: []v1.Server{
					*servers[0].Server,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			got, err := op.List(context.Background(), tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestServerOp_Read(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()

	tests := []struct {
		name     string
		serverId v1.ServerId
		want     *v1.Server
		wantErr  bool
	}{
		{
			name:     "minimum",
			serverId: v1.ServerId(servers[0].Id()),
			want:     servers[0].Server,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			got, err := op.Read(context.Background(), tt.serverId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestServerOp_ListOSImages(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()

	tests := []struct {
		name     string
		serverId v1.ServerId
		want     []*v1.OsImage
		wantErr  bool
	}{
		{
			name:     "minimum",
			serverId: v1.ServerId(servers[0].Id()),
			want:     servers[0].OSImages,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			got, err := op.ListOSImages(context.Background(), tt.serverId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListOSImages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestServerOp_OSInstall(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()

	type args struct {
		serverId v1.ServerId
		params   v1.OsInstallParameter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "minimum",
			args: args{
				serverId: v1.ServerId(servers[0].Id()),
				params: v1.OsInstallParameter{
					ManualPartition: true,
					OsImageId:       "usacloud",
					Password:        "passw0rd",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			if err := op.OSInstall(context.Background(), tt.args.serverId, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("OSInstall() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerOp_ReadPortChannel(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()

	type args struct {
		serverId      v1.ServerId
		portChannelId v1.PortChannelId
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.PortChannel
		wantErr bool
	}{
		{
			name: "minimum",
			args: args{
				serverId:      v1.ServerId(servers[0].Id()),
				portChannelId: 1001,
			},
			want:    &servers[0].Server.PortChannels[0],
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			got, err := op.ReadPortChannel(context.Background(), tt.args.serverId, tt.args.portChannelId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadPortChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestServerOp_ReadPort(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()

	type args struct {
		serverId v1.ServerId
		portId   v1.PortId
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.InterfacePort
		wantErr bool
	}{
		{
			name: "minimum",
			args: args{
				serverId: v1.ServerId(servers[0].Id()),
				portId:   2001,
			},
			want:    &servers[0].Server.Ports[0],
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			got, err := op.ReadPort(context.Background(), tt.args.serverId, tt.args.portId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestServerOp_UpdatePort(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()

	port := servers[0].Server.Ports[0]
	type args struct {
		serverId v1.ServerId
		portId   v1.PortId
		params   v1.UpdateServerPortParameter
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.InterfacePort
		wantErr bool
	}{
		{
			name: "minimum",
			args: args{
				serverId: v1.ServerId(servers[0].Id()),
				portId:   2001,
				params: v1.UpdateServerPortParameter{
					Nickname: "server01-port01-upd",
				},
			},
			want: &v1.InterfacePort{
				Enabled:             port.Enabled,
				GlobalBandwidthMbps: port.GlobalBandwidthMbps,
				Internet:            port.Internet,
				LocalBandwidthMbps:  port.LocalBandwidthMbps,
				Mode:                port.Mode,
				Nickname:            "server01-port01-upd",
				PortChannelId:       port.PortChannelId,
				PortId:              port.PortId,
				PrivateNetworks:     port.PrivateNetworks,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			got, err := op.UpdatePort(context.Background(), tt.args.serverId, tt.args.portId, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestServerOp_AssignNetwork(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()
	subnets := testServer.Engine.GetDedicatedSubnets()

	var internetType = v1.AssignNetworkParameterInternetTypeDedicatedSubnet
	var mode = v1.InterfacePortModeAccess
	port := servers[0].Server.Ports[0]

	type args struct {
		serverId v1.ServerId
		portId   v1.PortId
		params   v1.AssignNetworkParameter
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.InterfacePort
		wantErr bool
	}{
		{
			name: "minimum",
			args: args{
				serverId: v1.ServerId(servers[0].Id()),
				portId:   v1.PortId(port.PortId),
				params: v1.AssignNetworkParameter{
					DedicatedSubnetId: pointer.String(subnets[0].DedicatedSubnetId),
					InternetType:      &internetType,
					Mode:              v1.AssignNetworkParameterModeAccess,
				},
			},
			want: &v1.InterfacePort{
				Enabled:             port.Enabled,
				GlobalBandwidthMbps: pointer.Int(500),
				Internet: &v1.Internet{
					DedicatedSubnet: &v1.AttachedDedicatedSubnet{
						DedicatedSubnetId: subnets[0].DedicatedSubnetId,
						Nickname:          subnets[0].Service.Nickname,
					},
					NetworkAddress: subnets[0].Ipv4.NetworkAddress,
					PrefixLength:   subnets[0].Ipv4.PrefixLength,
					SubnetType:     v1.InternetSubnetTypeDedicatedSubnet,
				},
				PrivateNetworks:    nil,
				LocalBandwidthMbps: nil,
				Mode:               &mode,
				Nickname:           port.Nickname,
				PortChannelId:      port.PortChannelId,
				PortId:             port.PortId,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			got, err := op.AssignNetwork(context.Background(), tt.args.serverId, tt.args.portId, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssignNetwork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestServerOp_EnablePort(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()

	port := servers[0].Server.Ports[0]
	type args struct {
		serverId v1.ServerId
		portId   v1.PortId
		enable   bool
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.InterfacePort
		wantErr bool
	}{
		{
			name: "minimum",
			args: args{
				serverId: v1.ServerId(servers[0].Id()),
				portId:   v1.PortId(port.PortId),
				enable:   false,
			},
			want: &v1.InterfacePort{
				Enabled:             false,
				GlobalBandwidthMbps: port.GlobalBandwidthMbps,
				Internet:            port.Internet,
				LocalBandwidthMbps:  port.LocalBandwidthMbps,
				Mode:                port.Mode,
				Nickname:            port.Nickname,
				PortChannelId:       port.PortChannelId,
				PortId:              port.PortId,
				PrivateNetworks:     port.PrivateNetworks,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			got, err := op.EnablePort(context.Background(), tt.args.serverId, tt.args.portId, tt.args.enable)
			if (err != nil) != tt.wantErr {
				t.Errorf("EnablePort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestServerOp_ReadTrafficByPort(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()

	port := servers[0].Server.Ports[0]
	type args struct {
		serverId v1.ServerId
		portId   v1.PortId
		params   v1.ReadServerTrafficByPortParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "minimum",
			args: args{
				serverId: v1.ServerId(servers[0].Id()),
				portId:   v1.PortId(port.PortId),
				params:   v1.ReadServerTrafficByPortParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			got, err := op.ReadTrafficByPort(context.Background(), tt.args.serverId, tt.args.portId, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadTrafficByPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.NotNil(t, got)
			require.NotEmpty(t, got.Receive)
			require.NotEmpty(t, got.Transmit)
		})
	}
}

func TestServerOp_PowerControl(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()

	type args struct {
		serverId  v1.ServerId
		operation v1.ServerPowerOperations
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "minimum",
			args: args{
				serverId:  v1.ServerId(servers[0].Id()),
				operation: v1.ServerPowerOperationsReset,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			if err := op.PowerControl(context.Background(), tt.args.serverId, tt.args.operation); (err != nil) != tt.wantErr {
				t.Errorf("PowerControl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerOp_ReadPowerStatus(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()

	tests := []struct {
		name     string
		serverId v1.ServerId
		want     *v1.ServerPowerStatus
		wantErr  bool
	}{
		{
			name:     "minimum",
			serverId: v1.ServerId(servers[0].Id()),
			want:     servers[0].PowerStatus,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			got, err := op.ReadPowerStatus(context.Background(), tt.serverId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadPowerStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestServerOp_ReadRAIDStatus(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()

	type args struct {
		serverId v1.ServerId
		refresh  bool
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.RaidStatus
		wantErr bool
	}{
		{
			name: "minimum",
			args: args{
				serverId: v1.ServerId(servers[0].Id()),
				refresh:  false,
			},
			want:    servers[0].RaidStatus,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			got, err := op.ReadRAIDStatus(context.Background(), tt.args.serverId, tt.args.refresh)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadRAIDStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestServerOp_ConfigureBonding(t *testing.T) {
	onlyUnitTest(t)
	servers := testServer.Engine.GetServers()
	portChannel := servers[0].Server.PortChannels[0]

	type args struct {
		serverId      v1.ServerId
		portChannelId v1.PortChannelId
		params        v1.ConfigureBondingParameter
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.PortChannel
		wantErr bool
	}{
		{
			name: "minimum",
			args: args{
				serverId:      v1.ServerId(servers[0].Id()),
				portChannelId: v1.PortChannelId(portChannel.PortChannelId),
				params: v1.ConfigureBondingParameter{
					BondingType:   v1.BondingTypeLacp,
					PortNicknames: pointer.StringSlice([]string{"port01"}),
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServerOp{
				client: testClient(t),
			}
			got, err := op.ConfigureBonding(context.Background(), tt.args.serverId, tt.args.portChannelId, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigureBonding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.Equal(t, portChannel.PortChannelId, got.PortChannelId)
			require.NotEqualValues(t, portChannel.Ports, got.Ports)
			switch tt.args.params.BondingType {
			case v1.BondingTypeSingle:
				require.Len(t, got.Ports, 2)
			default:
				require.Len(t, got.Ports, 1)
			}
		})
	}
}
