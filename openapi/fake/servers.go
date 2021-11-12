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

// ListServers サーバー一覧
// (GET /servers/)
func (s *Server) ListServers(c *gin.Context, params openapi.ListServersParams) {
}

// ReadServer サーバー
// (GET /servers/{server_id}/)
func (s *Server) ReadServer(c *gin.Context, serverId openapi.ServerId) {
}

// ListOSImages インストール可能OS一覧
// (GET /servers/{server_id}/os_images/)
func (s *Server) ListOSImages(c *gin.Context, serverId openapi.ServerId) {
}

// OSInstall OSインストールの実行
// (POST /servers/{server_id}/os_install/)
func (s *Server) OSInstall(c *gin.Context, serverId openapi.ServerId, params openapi.OSInstallParams) {
}

// ReadServerPortChannel ポートチャネル状態取得
// (GET /servers/{server_id}/port_channels/{port_channel_id}/)
func (s *Server) ReadServerPortChannel(c *gin.Context, serverId openapi.ServerId, portChannelId openapi.PortChannelId) {
}

// SetPortChannelBonding ポートチャネル ボンディング設定
// (POST /servers/{server_id}/port_channels/{port_channel_id}/configure_bonding/)
func (s *Server) SetPortChannelBonding(c *gin.Context, serverId openapi.ServerId, portChannelId openapi.PortChannelId, params openapi.SetPortChannelBondingParams) {
}

// ReadServerPort ポート情報取得
// (GET /servers/{server_id}/ports/{port_id}/)
func (s *Server) ReadServerPort(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId) {
}

// UpdateServerPort ポート名称設定
// (PATCH /servers/{server_id}/ports/{port_id}/)
func (s *Server) UpdateServerPort(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.UpdateServerPortParams) {
}

// SetServerPortNetworkConnection ネットワーク接続設定の変更
// (POST /servers/{server_id}/ports/{port_id}/assign_network/)
func (s *Server) SetServerPortNetworkConnection(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.SetServerPortNetworkConnectionParams) {
}

// SetServerPortEnabled ポート有効/無効設定
// (POST /servers/{server_id}/ports/{port_id}/enable/)
func (s *Server) SetServerPortEnabled(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.SetServerPortEnabledParams) {
}

// ReadServerTrafficByPort トラフィックデータ取得
// (GET /servers/{server_id}/ports/{port_id}/traffic_graph/)
func (s *Server) ReadServerTrafficByPort(c *gin.Context, serverId openapi.ServerId, portId openapi.PortId, params openapi.ReadServerTrafficByPortParams) {
}

// SetServerPowerStatus サーバーの電源操作
// (POST /servers/{server_id}/power_control/)
func (s *Server) SetServerPowerStatus(c *gin.Context, serverId openapi.ServerId, params openapi.SetServerPowerStatusParams) {
}

// ReadServerPowerStatus サーバーの電源情報を取得する
// (GET /servers/{server_id}/power_status/)
func (s *Server) ReadServerPowerStatus(c *gin.Context, serverId openapi.ServerId) {
}

// ReadRAIDStatus サーバーのRAID状態を取得
// (GET /servers/{server_id}/raid_status/)
func (s *Server) ReadRAIDStatus(c *gin.Context, serverId openapi.ServerId, params openapi.ReadRAIDStatusParams) {
}
