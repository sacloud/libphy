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
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/sacloud/phy-go/openapi"
)

type Server struct {
	httpServer *httptest.Server

	// ListDedicatedSubnetsFunc 専用グローバルネットワーク 一覧
	// (GET /dedicated_subnets/)
	ListDedicatedSubnetsFunc func(c *gin.Context, params openapi.ListDedicatedSubnetsParams)

	// ReadDedicatedSubnetFunc 専用グローバルネットワーク
	// (GET /dedicated_subnets/{dedicated_subnet_id}/)
	ReadDedicatedSubnetFunc func(c *gin.Context, dedicatedSubnetId openapi.DedicatedSubnetId, params openapi.ReadDedicatedSubnetParams)

	// ListPrivateNetworksFunc ローカルネットワーク 一覧
	// (GET /private_networks/)
	ListPrivateNetworksFunc func(c *gin.Context, params openapi.ListPrivateNetworksParams)

	// ReadPrivateNetworkFunc ローカルネットワーク 詳細
	// (GET /private_networks/{private_network_id}/)
	ReadPrivateNetworkFunc func(c *gin.Context, privateNetworkId openapi.PrivateNetworkId)

	// ListServersFunc サーバー一覧
	// (GET /servers/)
	ListServersFunc func(c *gin.Context, params openapi.ListServersParams)

	// ReadServerFunc サーバー
	// (GET /servers/{server_id}/)
	ReadServerFunc func(c *gin.Context, serverId openapi.ServerId)

	// ListOSImagesFunc インストール可能OS一覧
	// (GET /servers/{server_id}/os_images/)
	ListOSImagesFunc func(c *gin.Context, serverId openapi.ServerId)

	// OSInstallFunc OSインストールの実行
	// (POST /servers/{server_id}/os_install/)
	OSInstallFunc func(c *gin.Context, serverId openapi.ServerId, params openapi.OSInstallParams)

	// ReadServerPortChannelFunc ポートチャネル状態取得
	// (GET /servers/{server_id}/port_channels/{port_channel_id}/)
	ReadServerPortChannelFunc func(c *gin.Context, serverId openapi.ServerId, portChannelId openapi.PortChannelId)

	// ServerConfigureBondingFunc ポートチャネル ボンディング設定
	// (POST /servers/{server_id}/port_channels/{port_channel_id}/configure_bonding/)
	ServerConfigureBondingFunc func(c *gin.Context, serverId openapi.ServerId, portChannelId openapi.PortChannelId, params openapi.ServerConfigureBondingParams)

	// ReadServerPortFunc ポート情報取得
	// (GET /servers/{server_id}/ports/{port_id}/)
	ReadServerPortFunc func(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId)

	// UpdateServerPortFunc ポート名称設定
	// (PATCH /servers/{server_id}/ports/{port_id}/)
	UpdateServerPortFunc func(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.UpdateServerPortParams)

	// ServerAssignNetworkFunc ネットワーク接続設定の変更
	// (POST /servers/{server_id}/ports/{port_id}/assign_network/)
	ServerAssignNetworkFunc func(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.ServerAssignNetworkParams)

	// EnableServerPortFunc ポート有効/無効設定
	// (POST /servers/{server_id}/ports/{port_id}/enable/)
	EnableServerPortFunc func(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.EnableServerPortParams)

	// ReadServerTrafficByPortFunc トラフィックデータ取得
	// (GET /servers/{server_id}/ports/{port_id}/traffic_graph/)
	ReadServerTrafficByPortFunc func(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.ReadServerTrafficByPortParams)

	// ServerPowerControlFunc サーバーの電源操作
	// (POST /servers/{server_id}/power_control/)
	ServerPowerControlFunc func(c *gin.Context, serverId openapi.ServerId, params openapi.ServerPowerControlParams)

	// ReadServerPowerStatusFunc サーバーの電源情報を取得する
	// (GET /servers/{server_id}/power_status/)
	ReadServerPowerStatusFunc func(c *gin.Context, serverId openapi.ServerId)

	// ReadRAIDStatusFunc サーバーのRAID状態を取得
	// (GET /servers/{server_id}/raid_status/)
	ReadRAIDStatusFunc func(c *gin.Context, serverId openapi.ServerId, params openapi.ReadRAIDStatusParams)

	// ListServicesFunc サービス一覧
	// (GET /services/)
	ListServicesFunc func(c *gin.Context, params openapi.ListServicesParams)

	// ReadServiceFunc サービス 詳細
	// (GET /services/{service_id}/)
	ReadServiceFunc func(c *gin.Context, serviceId openapi.ServiceId)

	// UpdateServiceFunc サービスの名称・説明の変更
	// (PATCH /services/{service_id}/)
	UpdateServiceFunc func(c *gin.Context, serviceId openapi.ServiceId, params openapi.UpdateServiceParams)
}

// Start テスト用のHTTPサーバを起動し、クリーンアップ用のfuncを返す
func (s *Server) Start() func() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	s.httpServer = httptest.NewServer(openapi.RegisterHandlers(router, s))
	return s.httpServer.Close
}

// URL テスト用サーバのURLを返す
func (s *Server) URL() string {
	if s.httpServer != nil {
		return s.httpServer.URL
	}
	panic("http server is not started")
}

// ListDedicatedSubnets 専用グローバルネットワーク 一覧
// (GET /dedicated_subnets/)
func (s *Server) ListDedicatedSubnets(c *gin.Context, params openapi.ListDedicatedSubnetsParams) {
	if s.ListDedicatedSubnetsFunc != nil {
		s.ListDedicatedSubnetsFunc(c, params)
	}
}

// ReadDedicatedSubnet 専用グローバルネットワーク
// (GET /dedicated_subnets/{dedicated_subnet_id}/)
func (s *Server) ReadDedicatedSubnet(c *gin.Context, dedicatedSubnetId openapi.DedicatedSubnetId, params openapi.ReadDedicatedSubnetParams) {
	if s.ReadDedicatedSubnetFunc != nil {
		s.ReadDedicatedSubnetFunc(c, dedicatedSubnetId, params)
	}
}

// ListPrivateNetworks ローカルネットワーク 一覧
// (GET /private_networks/)
func (s *Server) ListPrivateNetworks(c *gin.Context, params openapi.ListPrivateNetworksParams) {
	if s.ListPrivateNetworksFunc != nil {
		s.ListPrivateNetworksFunc(c, params)
	}
}

// ReadPrivateNetwork ローカルネットワーク 詳細
// (GET /private_networks/{private_network_id}/)
func (s *Server) ReadPrivateNetwork(c *gin.Context, privateNetworkId openapi.PrivateNetworkId) {
	if s.ReadPrivateNetworkFunc != nil {
		s.ReadPrivateNetworkFunc(c, privateNetworkId)
	}
}

// ListServers サーバー一覧
// (GET /servers/)
func (s *Server) ListServers(c *gin.Context, params openapi.ListServersParams) {
	if s.ListServersFunc != nil {
		s.ListServersFunc(c, params)
	}
}

// ReadServer サーバー
// (GET /servers/{server_id}/)
func (s *Server) ReadServer(c *gin.Context, serverId openapi.ServerId) {
	if s.ReadServerFunc != nil {
		s.ReadServerFunc(c, serverId)
	}
}

// ListOSImages インストール可能OS一覧
// (GET /servers/{server_id}/os_images/)
func (s *Server) ListOSImages(c *gin.Context, serverId openapi.ServerId) {
	if s.ListOSImagesFunc != nil {
		s.ListOSImagesFunc(c, serverId)
	}
}

// OSInstall OSインストールの実行
// (POST /servers/{server_id}/os_install/)
func (s *Server) OSInstall(c *gin.Context, serverId openapi.ServerId, params openapi.OSInstallParams) {
	if s.OSInstallFunc != nil {
		s.OSInstallFunc(c, serverId, params)
	}
}

// ReadServerPortChannel ポートチャネル状態取得
// (GET /servers/{server_id}/port_channels/{port_channel_id}/)
func (s *Server) ReadServerPortChannel(c *gin.Context, serverId openapi.ServerId, portChannelId openapi.PortChannelId) {
	if s.ReadServerPortChannelFunc != nil {
		s.ReadServerPortChannelFunc(c, serverId, portChannelId)
	}
}

// ServerConfigureBonding ポートチャネル ボンディング設定
// (POST /servers/{server_id}/port_channels/{port_channel_id}/configure_bonding/)
func (s *Server) ServerConfigureBonding(c *gin.Context, serverId openapi.ServerId, portChannelId openapi.PortChannelId, params openapi.ServerConfigureBondingParams) {
	if s.ServerConfigureBondingFunc != nil {
		s.ServerConfigureBondingFunc(c, serverId, portChannelId, params)
	}
}

// ReadServerPort ポート情報取得
// (GET /servers/{server_id}/ports/{port_id}/)
func (s *Server) ReadServerPort(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId) {
	if s.ReadServerPortFunc != nil {
		s.ReadServerPortFunc(c, serverId, portId)
	}
}

// UpdateServerPort ポート名称設定
// (PATCH /servers/{server_id}/ports/{port_id}/)
func (s *Server) UpdateServerPort(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.UpdateServerPortParams) {
	if s.UpdateServerPortFunc != nil {
		s.UpdateServerPortFunc(c, serverId, portId, params)
	}
}

// ServerAssignNetwork ネットワーク接続設定の変更
// (POST /servers/{server_id}/ports/{port_id}/assign_network/)
func (s *Server) ServerAssignNetwork(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.ServerAssignNetworkParams) {
	if s.ServerAssignNetworkFunc != nil {
		s.ServerAssignNetworkFunc(c, serverId, portId, params)
	}
}

// EnableServerPort ポート有効/無効設定
// (POST /servers/{server_id}/ports/{port_id}/enable/)
func (s *Server) EnableServerPort(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.EnableServerPortParams) {
	if s.EnableServerPortFunc != nil {
		s.EnableServerPortFunc(c, serverId, portId, params)
	}
}

// ReadServerTrafficByPort トラフィックデータ取得
// (GET /servers/{server_id}/ports/{port_id}/traffic_graph/)
func (s *Server) ReadServerTrafficByPort(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.ReadServerTrafficByPortParams) {
	if s.ReadServerTrafficByPortFunc != nil {
		s.ReadServerTrafficByPortFunc(c, serverId, portId, params)
	}
}

// ServerPowerControl サーバーの電源操作
// (POST /servers/{server_id}/power_control/)
func (s *Server) ServerPowerControl(c *gin.Context, serverId openapi.ServerId, params openapi.ServerPowerControlParams) {
	if s.ServerPowerControlFunc != nil {
		s.ServerPowerControlFunc(c, serverId, params)
	}
}

// ReadServerPowerStatus サーバーの電源情報を取得する
// (GET /servers/{server_id}/power_status/)
func (s *Server) ReadServerPowerStatus(c *gin.Context, serverId openapi.ServerId) {
	if s.ReadServerPowerStatusFunc != nil {
		s.ReadServerPowerStatusFunc(c, serverId)
	}
}

// ReadRAIDStatus サーバーのRAID状態を取得
// (GET /servers/{server_id}/raid_status/)
func (s *Server) ReadRAIDStatus(c *gin.Context, serverId openapi.ServerId, params openapi.ReadRAIDStatusParams) {
	if s.ReadRAIDStatusFunc != nil {
		s.ReadRAIDStatusFunc(c, serverId, params)
	}
}

// ListServices サービス一覧
// (GET /services/)
func (s *Server) ListServices(c *gin.Context, params openapi.ListServicesParams) {
	if s.ListServicesFunc != nil {
		s.ListServicesFunc(c, params)
	}
}

// ReadService サービス 詳細
// (GET /services/{service_id}/)
func (s *Server) ReadService(c *gin.Context, serviceId openapi.ServiceId) {
	if s.ReadServiceFunc != nil {
		s.ReadServiceFunc(c, serviceId)
	}
}

// UpdateService サービスの名称・説明の変更
// (PATCH /services/{service_id}/)
func (s *Server) UpdateService(c *gin.Context, serviceId openapi.ServiceId, params openapi.UpdateServiceParams) {
	if s.UpdateServiceFunc != nil {
		s.UpdateServiceFunc(c, serviceId, params)
	}
}
