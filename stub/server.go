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

package stub

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

type Server struct {
	// ListDedicatedSubnetsFunc 専用グローバルネットワーク 一覧
	// (GET /dedicated_subnets/)
	ListDedicatedSubnetsFunc func(c *gin.Context, params v1.ListDedicatedSubnetsParams)

	// ReadDedicatedSubnetFunc 専用グローバルネットワーク
	// (GET /dedicated_subnets/{dedicated_subnet_id}/)
	ReadDedicatedSubnetFunc func(c *gin.Context, dedicatedSubnetId v1.DedicatedSubnetId, params v1.ReadDedicatedSubnetParams)

	// ListPrivateNetworksFunc ローカルネットワーク 一覧
	// (GET /private_networks/)
	ListPrivateNetworksFunc func(c *gin.Context, params v1.ListPrivateNetworksParams)

	// ReadPrivateNetworkFunc ローカルネットワーク 詳細
	// (GET /private_networks/{private_network_id}/)
	ReadPrivateNetworkFunc func(c *gin.Context, privateNetworkId v1.PrivateNetworkId)

	// ListServersFunc サーバー一覧
	// (GET /servers/)
	ListServersFunc func(c *gin.Context, params v1.ListServersParams)

	// ReadServerFunc サーバー
	// (GET /servers/{server_id}/)
	ReadServerFunc func(c *gin.Context, serverId v1.ServerId)

	// ListOSImagesFunc インストール可能OS一覧
	// (GET /servers/{server_id}/os_images/)
	ListOSImagesFunc func(c *gin.Context, serverId v1.ServerId)

	// OSInstallFunc OSインストールの実行
	// (POST /servers/{server_id}/os_install/)
	OSInstallFunc func(c *gin.Context, serverId v1.ServerId, params v1.OSInstallParams)

	// ReadServerPortChannelFunc ポートチャネル状態取得
	// (GET /servers/{server_id}/port_channels/{port_channel_id}/)
	ReadServerPortChannelFunc func(c *gin.Context, serverId v1.ServerId, portChannelId v1.PortChannelId)

	// ServerConfigureBondingFunc ポートチャネル ボンディング設定
	// (POST /servers/{server_id}/port_channels/{port_channel_id}/configure_bonding/)
	ServerConfigureBondingFunc func(c *gin.Context, serverId v1.ServerId, portChannelId v1.PortChannelId, params v1.ServerConfigureBondingParams)

	// ReadServerPortFunc ポート情報取得
	// (GET /servers/{server_id}/ports/{port_id}/)
	ReadServerPortFunc func(c *gin.Context, serverId v1.ServerId, portId v1.PortId)

	// UpdateServerPortFunc ポート名称設定
	// (PATCH /servers/{server_id}/ports/{port_id}/)
	UpdateServerPortFunc func(c *gin.Context, serverId v1.ServerId, portId v1.PortId, params v1.UpdateServerPortParams)

	// ServerAssignNetworkFunc ネットワーク接続設定の変更
	// (POST /servers/{server_id}/ports/{port_id}/assign_network/)
	ServerAssignNetworkFunc func(c *gin.Context, serverId v1.ServerId, portId v1.PortId, params v1.ServerAssignNetworkParams)

	// EnableServerPortFunc ポート有効/無効設定
	// (POST /servers/{server_id}/ports/{port_id}/enable/)
	EnableServerPortFunc func(c *gin.Context, serverId v1.ServerId, portId v1.PortId, params v1.EnableServerPortParams)

	// ReadServerTrafficByPortFunc トラフィックデータ取得
	// (GET /servers/{server_id}/ports/{port_id}/traffic_graph/)
	ReadServerTrafficByPortFunc func(c *gin.Context, serverId v1.ServerId, portId v1.PortId, params v1.ReadServerTrafficByPortParams)

	// ServerPowerControlFunc サーバーの電源操作
	// (POST /servers/{server_id}/power_control/)
	ServerPowerControlFunc func(c *gin.Context, serverId v1.ServerId, params v1.ServerPowerControlParams)

	// ReadServerPowerStatusFunc サーバーの電源情報を取得する
	// (GET /servers/{server_id}/power_status/)
	ReadServerPowerStatusFunc func(c *gin.Context, serverId v1.ServerId)

	// ReadRAIDStatusFunc サーバーのRAID状態を取得
	// (GET /servers/{server_id}/raid_status/)
	ReadRAIDStatusFunc func(c *gin.Context, serverId v1.ServerId, params v1.ReadRAIDStatusParams)

	// ListServicesFunc サービス一覧
	// (GET /services/)
	ListServicesFunc func(c *gin.Context, params v1.ListServicesParams)

	// ReadServiceFunc サービス 詳細
	// (GET /services/{service_id}/)
	ReadServiceFunc func(c *gin.Context, serviceId v1.ServiceId)

	// UpdateServiceFunc サービスの名称・説明の変更
	// (PATCH /services/{service_id}/)
	UpdateServiceFunc func(c *gin.Context, serviceId v1.ServiceId, params v1.UpdateServiceParams)
}

// Handler 構築済みのhttp.Handlerを返す
func (s *Server) Handler() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	if os.Getenv("PHY_SERVER_DEBUG") != "" {
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	if os.Getenv("PHY_SERVER_LOGGING") != "" {
		engine.Use(gin.Logger())
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return v1.RegisterHandlers(engine, s)
}

// ListDedicatedSubnets 専用グローバルネットワーク 一覧
// (GET /dedicated_subnets/)
func (s *Server) ListDedicatedSubnets(c *gin.Context, params v1.ListDedicatedSubnetsParams) {
	if s.ListDedicatedSubnetsFunc != nil {
		s.ListDedicatedSubnetsFunc(c, params)
	}
}

// ReadDedicatedSubnet 専用グローバルネットワーク
// (GET /dedicated_subnets/{dedicated_subnet_id}/)
func (s *Server) ReadDedicatedSubnet(c *gin.Context, dedicatedSubnetId v1.DedicatedSubnetId, params v1.ReadDedicatedSubnetParams) {
	if s.ReadDedicatedSubnetFunc != nil {
		s.ReadDedicatedSubnetFunc(c, dedicatedSubnetId, params)
	}
}

// ListPrivateNetworks ローカルネットワーク 一覧
// (GET /private_networks/)
func (s *Server) ListPrivateNetworks(c *gin.Context, params v1.ListPrivateNetworksParams) {
	if s.ListPrivateNetworksFunc != nil {
		s.ListPrivateNetworksFunc(c, params)
	}
}

// ReadPrivateNetwork ローカルネットワーク 詳細
// (GET /private_networks/{private_network_id}/)
func (s *Server) ReadPrivateNetwork(c *gin.Context, privateNetworkId v1.PrivateNetworkId) {
	if s.ReadPrivateNetworkFunc != nil {
		s.ReadPrivateNetworkFunc(c, privateNetworkId)
	}
}

// ListServers サーバー一覧
// (GET /servers/)
func (s *Server) ListServers(c *gin.Context, params v1.ListServersParams) {
	if s.ListServersFunc != nil {
		s.ListServersFunc(c, params)
	}
}

// ReadServer サーバー
// (GET /servers/{server_id}/)
func (s *Server) ReadServer(c *gin.Context, serverId v1.ServerId) {
	if s.ReadServerFunc != nil {
		s.ReadServerFunc(c, serverId)
	}
}

// ListOSImages インストール可能OS一覧
// (GET /servers/{server_id}/os_images/)
func (s *Server) ListOSImages(c *gin.Context, serverId v1.ServerId) {
	if s.ListOSImagesFunc != nil {
		s.ListOSImagesFunc(c, serverId)
	}
}

// OSInstall OSインストールの実行
// (POST /servers/{server_id}/os_install/)
func (s *Server) OSInstall(c *gin.Context, serverId v1.ServerId, params v1.OSInstallParams) {
	if s.OSInstallFunc != nil {
		s.OSInstallFunc(c, serverId, params)
	}
}

// ReadServerPortChannel ポートチャネル状態取得
// (GET /servers/{server_id}/port_channels/{port_channel_id}/)
func (s *Server) ReadServerPortChannel(c *gin.Context, serverId v1.ServerId, portChannelId v1.PortChannelId) {
	if s.ReadServerPortChannelFunc != nil {
		s.ReadServerPortChannelFunc(c, serverId, portChannelId)
	}
}

// ServerConfigureBonding ポートチャネル ボンディング設定
// (POST /servers/{server_id}/port_channels/{port_channel_id}/configure_bonding/)
func (s *Server) ServerConfigureBonding(c *gin.Context, serverId v1.ServerId, portChannelId v1.PortChannelId, params v1.ServerConfigureBondingParams) {
	if s.ServerConfigureBondingFunc != nil {
		s.ServerConfigureBondingFunc(c, serverId, portChannelId, params)
	}
}

// ReadServerPort ポート情報取得
// (GET /servers/{server_id}/ports/{port_id}/)
func (s *Server) ReadServerPort(c *gin.Context, serverId v1.ServerId, portId v1.PortId) {
	if s.ReadServerPortFunc != nil {
		s.ReadServerPortFunc(c, serverId, portId)
	}
}

// UpdateServerPort ポート名称設定
// (PATCH /servers/{server_id}/ports/{port_id}/)
func (s *Server) UpdateServerPort(c *gin.Context, serverId v1.ServerId, portId v1.PortId, params v1.UpdateServerPortParams) {
	if s.UpdateServerPortFunc != nil {
		s.UpdateServerPortFunc(c, serverId, portId, params)
	}
}

// ServerAssignNetwork ネットワーク接続設定の変更
// (POST /servers/{server_id}/ports/{port_id}/assign_network/)
func (s *Server) ServerAssignNetwork(c *gin.Context, serverId v1.ServerId, portId v1.PortId, params v1.ServerAssignNetworkParams) {
	if s.ServerAssignNetworkFunc != nil {
		s.ServerAssignNetworkFunc(c, serverId, portId, params)
	}
}

// EnableServerPort ポート有効/無効設定
// (POST /servers/{server_id}/ports/{port_id}/enable/)
func (s *Server) EnableServerPort(c *gin.Context, serverId v1.ServerId, portId v1.PortId, params v1.EnableServerPortParams) {
	if s.EnableServerPortFunc != nil {
		s.EnableServerPortFunc(c, serverId, portId, params)
	}
}

// ReadServerTrafficByPort トラフィックデータ取得
// (GET /servers/{server_id}/ports/{port_id}/traffic_graph/)
func (s *Server) ReadServerTrafficByPort(c *gin.Context, serverId v1.ServerId, portId v1.PortId, params v1.ReadServerTrafficByPortParams) {
	if s.ReadServerTrafficByPortFunc != nil {
		s.ReadServerTrafficByPortFunc(c, serverId, portId, params)
	}
}

// ServerPowerControl サーバーの電源操作
// (POST /servers/{server_id}/power_control/)
func (s *Server) ServerPowerControl(c *gin.Context, serverId v1.ServerId, params v1.ServerPowerControlParams) {
	if s.ServerPowerControlFunc != nil {
		s.ServerPowerControlFunc(c, serverId, params)
	}
}

// ReadServerPowerStatus サーバーの電源情報を取得する
// (GET /servers/{server_id}/power_status/)
func (s *Server) ReadServerPowerStatus(c *gin.Context, serverId v1.ServerId) {
	if s.ReadServerPowerStatusFunc != nil {
		s.ReadServerPowerStatusFunc(c, serverId)
	}
}

// ReadRAIDStatus サーバーのRAID状態を取得
// (GET /servers/{server_id}/raid_status/)
func (s *Server) ReadRAIDStatus(c *gin.Context, serverId v1.ServerId, params v1.ReadRAIDStatusParams) {
	if s.ReadRAIDStatusFunc != nil {
		s.ReadRAIDStatusFunc(c, serverId, params)
	}
}

// ListServices サービス一覧
// (GET /services/)
func (s *Server) ListServices(c *gin.Context, params v1.ListServicesParams) {
	if s.ListServicesFunc != nil {
		s.ListServicesFunc(c, params)
	}
}

// ReadService サービス 詳細
// (GET /services/{service_id}/)
func (s *Server) ReadService(c *gin.Context, serviceId v1.ServiceId) {
	if s.ReadServiceFunc != nil {
		s.ReadServiceFunc(c, serviceId)
	}
}

// UpdateService サービスの名称・説明の変更
// (PATCH /services/{service_id}/)
func (s *Server) UpdateService(c *gin.Context, serviceId v1.ServiceId, params v1.UpdateServiceParams) {
	if s.UpdateServiceFunc != nil {
		s.UpdateServiceFunc(c, serviceId, params)
	}
}
