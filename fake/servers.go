// Copyright 2021-2022 The phy-go authors
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
	"fmt"
	"time"

	"github.com/getlantern/deepcopy"
	v1 "github.com/sacloud/phy-go/apis/v1"
)

// ListServers サーバー一覧
// (GET /servers/)
func (engine *Engine) ListServers(params v1.ListServersParams) (*v1.Servers, error) {
	defer engine.rLock()()

	// TODO 検索条件の処理を実装

	return &v1.Servers{
		Meta: v1.PaginateMeta{
			Count: len(engine.Servers),
		},
		Servers: engine.servers(),
	}, nil
}

// ReadServer サーバー
// (GET /servers/{server_id}/)
func (engine *Engine) ReadServer(serverId v1.ServerId) (*v1.Server, error) {
	defer engine.rLock()()

	s := engine.getServerById(serverId)
	if s != nil {
		var server v1.Server
		if err := deepcopy.Copy(&server, s.Server); err != nil {
			return nil, err
		}
		return &server, nil
	}
	return nil, NewError(ErrorTypeNotFound, "server", string(serverId))
}

// ListOSImages インストール可能OS一覧
// (GET /servers/{server_id}/os_images/)
func (engine *Engine) ListOSImages(serverId v1.ServerId) ([]*v1.OsImage, error) {
	defer engine.rLock()()

	s := engine.getServerById(serverId)
	if s != nil {
		var images []*v1.OsImage
		if err := deepcopy.Copy(&images, s.OSImages); err != nil {
			return nil, err
		}
		return images, nil
	}
	return nil, NewError(ErrorTypeNotFound, "server", string(serverId))
}

// OSInstall OSインストールの実行
// (POST /servers/{server_id}/os_install/)
func (engine *Engine) OSInstall(serverId v1.ServerId, params v1.OsInstallParameter) error {
	defer engine.rLock()()

	s := engine.getServerById(serverId)
	if s != nil {
		if s.Server.LockStatus != nil {
			return NewError(ErrorTypeConflict, "server", string(serverId))
		}
		for _, image := range s.OSImages {
			if image.OsImageId == params.OsImageId {
				engine.startOSInstall(s)
				return nil
			}
		}
		return NewError(ErrorTypeNotFound, "os-image", params.OsImageId, "server[%s]", serverId)
	}
	return NewError(ErrorTypeNotFound, "server", serverId)
}

// ReadServerPortChannel ポートチャネル状態取得
// (GET /servers/{server_id}/port_channels/{port_channel_id}/)
func (engine *Engine) ReadServerPortChannel(serverId v1.ServerId, portChannelId v1.PortChannelId) (*v1.PortChannel, error) {
	defer engine.rLock()()

	s := engine.getServerById(serverId)
	if s != nil {
		return s.getPortChannelById(portChannelId)
	}
	return nil, NewError(ErrorTypeNotFound, "server", serverId)
}

// ServerConfigureBonding ポートチャネル ボンディング設定
// (POST /servers/{server_id}/port_channels/{port_channel_id}/configure_bonding/)
//
// この実装では排他ロックをかけて同期的に処理するため対象ポートチャネルのLockedは更新しない
func (engine *Engine) ServerConfigureBonding(serverId v1.ServerId, portChannelId v1.PortChannelId, params v1.ConfigureBondingParameter) (*v1.PortChannel, error) {
	defer engine.lock()() // ここで同期的に更新処理を行うため書き込みロック

	s := engine.getServerById(serverId)
	if s != nil {
		portChannel, err := s.getPortChannelById(portChannelId)
		if err != nil {
			return nil, err
		}
		// BondingTypeとserver.spec.{port_channel_1gbe_count, port_channel_10gbe_count}に応じて
		// 必要な数だけportを作成
		var ports []v1.InterfacePort
		var portIds []int

		switch params.BondingType {
		case v1.BondingTypeLacp, v1.BondingTypeStatic:
			if params.PortNicknames != nil && len(*params.PortNicknames) != 1 {
				return nil, NewError(ErrorTypeInvalidRequest, "port-channel", portChannelId, "invalid PortNicknames")
			}
			name := string(portChannel.LinkSpeedType)
			if params.PortNicknames != nil {
				names := *params.PortNicknames
				if names[0] != "" {
					name = names[0]
				}
			}
			port := v1.InterfacePort{
				Enabled:       true,
				Nickname:      name,
				PortChannelId: portChannel.PortChannelId,
				PortId:        engine.nextId(),
			}

			ports = append(ports, port)
			portIds = append(portIds, port.PortId)
		case v1.BondingTypeSingle:
			if params.PortNicknames != nil && len(*params.PortNicknames) != 2 {
				return nil, NewError(ErrorTypeInvalidRequest, "port-channel", portChannelId, "invalid PortNicknames")
			}

			prefix := string(portChannel.LinkSpeedType)
			names := []string{prefix + " 1", prefix + " 2"}
			if params.PortNicknames != nil {
				names = *params.PortNicknames
			}
			for _, name := range names {
				port := v1.InterfacePort{
					Enabled:       true,
					Nickname:      name,
					PortChannelId: portChannel.PortChannelId,
					PortId:        engine.nextId(),
				}
				ports = append(ports, port)
				portIds = append(portIds, port.PortId)
			}
		}
		s.Server.Ports = ports
		portChannel.Ports = portIds

		s.updatePortChannel(portChannel)
		return portChannel, nil
	}
	return nil, NewError(ErrorTypeNotFound, "server", serverId)
}

// ReadServerPort ポート情報取得
// (GET /servers/{server_id}/ports/{port_id}/)
func (engine *Engine) ReadServerPort(serverId v1.ServerId, portId v1.PortId) (*v1.InterfacePort, error) {
	defer engine.rLock()()

	s := engine.getServerById(serverId)
	if s != nil {
		return s.getPortById(portId)
	}
	return nil, NewError(ErrorTypeNotFound, "server", serverId)
}

// UpdateServerPort ポート名称設定
// (PATCH /servers/{server_id}/ports/{port_id}/)
func (engine *Engine) UpdateServerPort(serverId v1.ServerId, portId v1.PortId, params v1.UpdateServerPortParameter) (*v1.InterfacePort, error) {
	defer engine.lock()()

	s := engine.getServerById(serverId)
	if s != nil {
		port, err := s.getPortById(portId)
		if err != nil {
			return nil, err
		}
		port.Nickname = params.Nickname

		s.updatePort(port)
		return port, nil
	}
	return nil, NewError(ErrorTypeNotFound, "server", serverId)
}

// ServerAssignNetwork ネットワーク接続設定の変更
// (POST /servers/{server_id}/ports/{port_id}/assign_network/)
//
// Note: この実装では本来不可能な複数のインターネット接続(がされたポート)を許容している。
// 必要に応じて利用者側で適切にハンドリングすること。
func (engine *Engine) ServerAssignNetwork(serverId v1.ServerId, portId v1.PortId, params v1.AssignNetworkParameter) (*v1.InterfacePort, error) {
	defer engine.lock()()

	s := engine.getServerById(serverId)
	if s != nil {
		port, err := s.getPortById(portId)
		if err != nil {
			return nil, err
		}

		// 一旦関連する項目をリセット
		port.Internet = nil
		port.Mode = nil
		port.PrivateNetworks = nil
		port.GlobalBandwidthMbps = nil
		port.LocalBandwidthMbps = nil

		var internet *v1.Internet
		if params.InternetType != nil {
			switch *params.InternetType {
			case v1.AssignNetworkParameterInternetTypeCommonSubnet:
				// TODO 共用グローバルネットをどこかに定義しておく
				internet = &v1.Internet{
					NetworkAddress: "203.0.113.0",
					PrefixLength:   24,
					SubnetType:     v1.InternetSubnetTypeCommonSubnet,
				}
				mbps := 100
				port.GlobalBandwidthMbps = &mbps
			case v1.AssignNetworkParameterInternetTypeDedicatedSubnet:
				subnet := engine.getDedicatedSubnetById(v1.DedicatedSubnetId(*params.DedicatedSubnetId))
				if subnet == nil {
					return nil, NewError(ErrorTypeInvalidRequest, "port", portId, "invalid dedicated subnet id: %s", params.DedicatedSubnetId)
				}
				internet = &v1.Internet{
					DedicatedSubnet: &v1.AttachedDedicatedSubnet{
						DedicatedSubnetId: subnet.DedicatedSubnetId,
						Nickname:          subnet.Service.Nickname,
					},
					NetworkAddress: subnet.Ipv4.NetworkAddress,
					PrefixLength:   subnet.Ipv4.PrefixLength,
					SubnetType:     v1.InternetSubnetTypeDedicatedSubnet,
				}
				mbps := 500
				port.GlobalBandwidthMbps = &mbps
			default:
				panic(fmt.Errorf("invalid InternetType: %v", params.InternetType))
			}
		}
		port.Internet = internet

		switch params.Mode {
		case v1.AssignNetworkParameterModeAccess:
			v := v1.InterfacePortModeAccess
			port.Mode = &v
		case v1.AssignNetworkParameterModeTrunk:
			v := v1.InterfacePortModeTrunk
			port.Mode = &v
		}

		if params.PrivateNetworkIds != nil {
			for _, id := range *params.PrivateNetworkIds {
				pn := engine.getPrivateNetworkById(v1.PrivateNetworkId(id))
				if pn == nil {
					return nil, NewError(ErrorTypeInvalidRequest, "port", portId, "invalid private network id: %s", id)
				}
				port.PrivateNetworks = append(port.PrivateNetworks, v1.AttachedPrivateNetwork{
					Nickname:         pn.Service.Nickname,
					PrivateNetworkId: pn.PrivateNetworkId,
				})
			}
			// 2000もあり得るがこの実装では1000で固定
			mbps := 1000
			port.LocalBandwidthMbps = &mbps
		}

		s.updatePort(port)
		return port, nil
	}
	return nil, NewError(ErrorTypeNotFound, "server", serverId)
}

// EnableServerPort ポート有効/無効設定
// (POST /servers/{server_id}/ports/{port_id}/enable/)
func (engine *Engine) EnableServerPort(serverId v1.ServerId, portId v1.PortId, params v1.EnableServerPortParameter) (*v1.InterfacePort, error) {
	defer engine.lock()()

	s := engine.getServerById(serverId)
	if s != nil {
		port, err := s.getPortById(portId)
		if err != nil {
			return nil, err
		}
		port.Enabled = params.Enable

		s.updatePort(port)
		return port, nil
	}
	return nil, NewError(ErrorTypeNotFound, "server", serverId)
}

// ReadServerTrafficByPort トラフィックデータ取得
// (GET /servers/{server_id}/ports/{port_id}/traffic_graph/)
//
// Note: この実装では対象サーバが存在する場合は固定のレスポンスを返すのみ
func (engine *Engine) ReadServerTrafficByPort(serverId v1.ServerId, portId v1.PortId, params v1.ReadServerTrafficByPortParams) (*v1.TrafficGraph, error) {
	defer engine.rLock()()

	s := engine.getServerById(serverId)
	if s != nil {
		return &v1.TrafficGraph{
			Receive: []v1.TrafficGraphData{
				{
					Timestamp: time.Now(),
					Value:     1,
				},
				{
					Timestamp: time.Now().Add(-1 * time.Minute),
					Value:     2,
				},
			},
			Transmit: []v1.TrafficGraphData{
				{
					Timestamp: time.Now(),
					Value:     1,
				},
				{
					Timestamp: time.Now().Add(-1 * time.Minute),
					Value:     2,
				},
			},
		}, nil
	}
	return nil, NewError(ErrorTypeNotFound, "server", serverId)
}

// ServerPowerControl サーバーの電源操作
// (POST /servers/{server_id}/power_control/)
func (engine *Engine) ServerPowerControl(serverId v1.ServerId, params v1.PowerControlParameter) error {
	defer engine.rLock()()

	s := engine.getServerById(serverId)
	if s != nil {
		if s.Server.LockStatus != nil {
			return NewError(ErrorTypeConflict, "server", string(serverId))
		}

		engine.startServerPowerControl(s, params)
		return nil
	}
	return NewError(ErrorTypeNotFound, "server", serverId)
}

// ReadServerPowerStatus サーバーの電源情報を取得する
// (GET /servers/{server_id}/power_status/)
func (engine *Engine) ReadServerPowerStatus(serverId v1.ServerId) (*v1.ServerPowerStatus, error) {
	defer engine.rLock()()

	s := engine.getServerById(serverId)
	if s != nil {
		return s.PowerStatus, nil
	}
	return nil, NewError(ErrorTypeNotFound, "server", serverId)
}

// ReadRAIDStatus サーバーのRAID状態を取得
// (GET /servers/{server_id}/raid_status/)
//
// Note: この実装ではrefreshパラメータは無視される
func (engine *Engine) ReadRAIDStatus(serverId v1.ServerId, params v1.ReadRAIDStatusParams) (*v1.RaidStatus, error) {
	defer engine.rLock()()

	s := engine.getServerById(serverId)
	if s != nil {
		return s.RaidStatus, nil
	}
	return nil, NewError(ErrorTypeNotFound, "server", serverId)
}

// servers []*ServerDataから[]v1.Serverに変換して返す
func (engine *Engine) servers() []v1.Server {
	var results []v1.Server
	for _, s := range engine.Servers {
		results = append(results, *s.Server)
	}
	return results
}

func (engine *Engine) getServerById(serverId v1.ServerId) *Server {
	for _, s := range engine.Servers {
		if s.Id() == string(serverId) {
			return s
		}
	}
	return nil
}

func (engine *Engine) startOSInstall(server *Server) {
	go engine.startUpdateAction(func() {
		//start
		status := v1.ServerLockStatusOsInstall
		server.Server.LockStatus = &status

		// finish
		go engine.startUpdateAction(func() {
			server.Server.LockStatus = nil
		})
	})
}

func (engine *Engine) startServerPowerControl(server *Server, params v1.PowerControlParameter) {
	var powerStates v1.ServerPowerStatusStatus
	var cachedPowerStatus v1.CachedPowerStatusStatus
	switch string(params.Operation) {
	case "on", "reset":
		powerStates = v1.ServerPowerStatusStatusOn
		cachedPowerStatus = v1.CachedPowerStatusStatusOn
	case "soft", "off":
		powerStates = v1.ServerPowerStatusStatusOff
		cachedPowerStatus = v1.CachedPowerStatusStatusOff
	}

	go engine.startUpdateAction(func() {
		server.PowerStatus = &v1.ServerPowerStatus{
			Status: powerStates,
		}
		server.Server.CachedPowerStatus = &v1.CachedPowerStatus{
			Status: cachedPowerStatus,
			Stored: time.Now(),
		}
	})
}
