// Copyright 2021-2023 The phy-api-go authors
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
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

// ListServers サーバー一覧
// (GET /servers/)
func (s *Server) ListServers(c *gin.Context, params v1.ListServersParams) {
	servers, err := s.Engine.ListServers(params)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, servers)
}

// ReadServer サーバー
// (GET /servers/{server_id}/)
func (s *Server) ReadServer(c *gin.Context, serverId v1.ServerId) {
	server, err := s.Engine.ReadServer(serverId)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.ResponseBodyServer{
		Server: *server,
	})
}

// ListOSImages インストール可能OS一覧
// (GET /servers/{server_id}/os_images/)
func (s *Server) ListOSImages(c *gin.Context, serverId v1.ServerId) {
	results, err := s.Engine.ListOSImages(serverId)
	if err != nil {
		s.handleError(c, err)
		return
	}

	var images []v1.OsImage
	for _, v := range results {
		images = append(images, *v)
	}
	c.JSON(http.StatusOK, &v1.ResponseBodyOsImages{
		OsImages: images,
	})
}

// OSInstall OSインストールの実行
// (POST /servers/{server_id}/os_install/)
func (s *Server) OSInstall(c *gin.Context, serverId v1.ServerId, _ v1.OSInstallParams) {
	var paramJSON v1.OsInstallParameter
	if err := c.ShouldBindJSON(&paramJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.Engine.OSInstall(serverId, paramJSON); err != nil {
		s.handleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// ReadServerPortChannel ポートチャネル状態取得
// (GET /servers/{server_id}/port_channels/{port_channel_id}/)
func (s *Server) ReadServerPortChannel(c *gin.Context, serverId v1.ServerId, portChannelId v1.PortChannelId) {
	portChannel, err := s.Engine.ReadServerPortChannel(serverId, portChannelId)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.ResponseBodyPortChannel{
		PortChannel: *portChannel,
	})
}

// ServerConfigureBonding ポートチャネル ボンディング設定
// (POST /servers/{server_id}/port_channels/{port_channel_id}/configure_bonding/)
func (s *Server) ServerConfigureBonding(c *gin.Context, serverId v1.ServerId, portChannelId v1.PortChannelId, _ v1.ServerConfigureBondingParams) {
	var paramJSON v1.ConfigureBondingParameter
	if err := c.ShouldBindJSON(&paramJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	portChannel, err := s.Engine.ServerConfigureBonding(serverId, portChannelId, paramJSON)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.ResponseBodyPortChannel{
		PortChannel: *portChannel,
	})
}

// ReadServerPort ポート情報取得
// (GET /servers/{server_id}/ports/{port_id}/)
func (s *Server) ReadServerPort(c *gin.Context, serverId v1.ServerId, portId v1.PortId) {
	port, err := s.Engine.ReadServerPort(serverId, portId)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.ResponseBodyPort{
		Port: *port,
	})
}

// UpdateServerPort ポート名称設定
// (PATCH /servers/{server_id}/ports/{port_id}/)
func (s *Server) UpdateServerPort(c *gin.Context, serverId v1.ServerId, portId v1.PortId, _ v1.UpdateServerPortParams) {
	var paramJSON v1.UpdateServerPortParameter
	if err := c.ShouldBindJSON(&paramJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	port, err := s.Engine.UpdateServerPort(serverId, portId, paramJSON)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.ResponseBodyPort{
		Port: *port,
	})
}

// ServerAssignNetwork ネットワーク接続設定の変更
// (POST /servers/{server_id}/ports/{port_id}/assign_network/)
func (s *Server) ServerAssignNetwork(c *gin.Context, serverId v1.ServerId, portId v1.PortId, _ v1.ServerAssignNetworkParams) {
	var paramJSON v1.AssignNetworkParameter
	if err := c.ShouldBindJSON(&paramJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	port, err := s.Engine.ServerAssignNetwork(serverId, portId, paramJSON)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.ResponseBodyPort{
		Port: *port,
	})
}

// EnableServerPort ポート有効/無効設定
// (POST /servers/{server_id}/ports/{port_id}/enable/)
func (s *Server) EnableServerPort(c *gin.Context, serverId v1.ServerId, portId v1.PortId, _ v1.EnableServerPortParams) {
	var paramJSON v1.EnableServerPortParameter
	if err := c.ShouldBindJSON(&paramJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	port, err := s.Engine.EnableServerPort(serverId, portId, paramJSON)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, v1.ResponseBodyPort{
		Port: *port,
	})
}

// ReadServerTrafficByPort トラフィックデータ取得
// (GET /servers/{server_id}/ports/{port_id}/traffic_graph/)
func (s *Server) ReadServerTrafficByPort(c *gin.Context, serverId v1.ServerId, portId v1.PortId, params v1.ReadServerTrafficByPortParams) {
	traffic, err := s.Engine.ReadServerTrafficByPort(serverId, portId, params)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.ResponseBodyTrafficGraph{
		TrafficGraph: *traffic,
	})
}

// ServerPowerControl サーバーの電源操作
// (POST /servers/{server_id}/power_control/)
func (s *Server) ServerPowerControl(c *gin.Context, serverId v1.ServerId, _ v1.ServerPowerControlParams) {
	var paramJSON v1.PowerControlParameter
	if err := c.ShouldBindJSON(&paramJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.Engine.ServerPowerControl(serverId, paramJSON); err != nil {
		s.handleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// ReadServerPowerStatus サーバーの電源情報を取得する
// (GET /servers/{server_id}/power_status/)
func (s *Server) ReadServerPowerStatus(c *gin.Context, serverId v1.ServerId) {
	ps, err := s.Engine.ReadServerPowerStatus(serverId)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.ResponseBodyServerPowerStatus{
		PowerStatus: *ps,
	})
}

// ReadRAIDStatus サーバーのRAID状態を取得
// (GET /servers/{server_id}/raid_status/)
func (s *Server) ReadRAIDStatus(c *gin.Context, serverId v1.ServerId, params v1.ReadRAIDStatusParams) {
	raidStatus, err := s.Engine.ReadRAIDStatus(serverId, params)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.ResponseBodyRaidStatus{
		RaidStatus: *raidStatus,
	})
}
