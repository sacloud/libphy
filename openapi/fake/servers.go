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
	"github.com/gin-gonic/gin"
	"github.com/sacloud/phy-go/openapi"
)

// GetServers サーバー一覧
// (GET /servers/)
func (s *Server) GetServers(c *gin.Context, params openapi.GetServersParams) {
}

// GetServersServerId サーバー
// (GET /servers/{server_id}/)
func (s *Server) GetServersServerId(c *gin.Context, serverId openapi.ServerId) {

}

// GetServersServerIdOsImages インストール可能OS一覧
// (GET /servers/{server_id}/os_images/)
func (s *Server) GetServersServerIdOsImages(c *gin.Context, serverId openapi.ServerId) {

}

// PostServersServerIdOsInstall OSインストールの実行
// (POST /servers/{server_id}/os_install/)
func (s *Server) PostServersServerIdOsInstall(c *gin.Context, serverId openapi.ServerId, params openapi.PostServersServerIdOsInstallParams) {

}

// GetServersServerIdPortChannelsPortChannelId ポートチャネル状態取得
// (GET /servers/{server_id}/port_channels/{port_channel_id}/)
func (s *Server) GetServersServerIdPortChannelsPortChannelId(c *gin.Context, serverId openapi.ServerId, portChannelId openapi.PortChannelId) {

}

// PostServersServerIdPortChannelsPortChannelIdConfigureBonding ポートチャネル ボンディング設定
// (POST /servers/{server_id}/port_channels/{port_channel_id}/configure_bonding/)
func (s *Server) PostServersServerIdPortChannelsPortChannelIdConfigureBonding(c *gin.Context, serverId openapi.ServerId, portChannelId openapi.PortChannelId, params openapi.PostServersServerIdPortChannelsPortChannelIdConfigureBondingParams) {

}

// GetServersServerIdPortsPortId ポート情報取得
// (GET /servers/{server_id}/ports/{port_id}/)
func (s *Server) GetServersServerIdPortsPortId(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId) {

}

// PatchServersServerIdPortsPortId ポート名称設定
// (PATCH /servers/{server_id}/ports/{port_id}/)
func (s *Server) PatchServersServerIdPortsPortId(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.PatchServersServerIdPortsPortIdParams) {

}

// PostServersServerIdPortsPortIdAssignNetwork ネットワーク接続設定の変更
// (POST /servers/{server_id}/ports/{port_id}/assign_network/)
func (s *Server) PostServersServerIdPortsPortIdAssignNetwork(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.PostServersServerIdPortsPortIdAssignNetworkParams) {

}

// PostServersServerIdPortsPortIdEnable ポート有効/無効設定
// (POST /servers/{server_id}/ports/{port_id}/enable/)
func (s *Server) PostServersServerIdPortsPortIdEnable(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.PostServersServerIdPortsPortIdEnableParams) {

}

// GetServersServerIdPortsPortIdTrafficGraph トラフィックデータ取得
// (GET /servers/{server_id}/ports/{port_id}/traffic_graph/)
func (s *Server) GetServersServerIdPortsPortIdTrafficGraph(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.GetServersServerIdPortsPortIdTrafficGraphParams) {

}

// PostServersServerIdPowerControl サーバーの電源操作
// (POST /servers/{server_id}/power_control/)
func (s *Server) PostServersServerIdPowerControl(c *gin.Context, serverId openapi.ServerId, params openapi.PostServersServerIdPowerControlParams) {

}

// GetServersServerIdPowerStatus サーバーの電源情報を取得する
// (GET /servers/{server_id}/power_status/)
func (s *Server) GetServersServerIdPowerStatus(c *gin.Context, serverId openapi.ServerId) {

}

// GetServersServerIdRaidStatus サーバーのRAID状態を取得
// (GET /servers/{server_id}/raid_status/)
func (s *Server) GetServersServerIdRaidStatus(c *gin.Context, serverId openapi.ServerId, params openapi.GetServersServerIdRaidStatusParams) {

}
