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

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

// ListPrivateNetworks ローカルネットワーク 一覧
// (GET /private_networks/)
func (s *Server) ListPrivateNetworks(c *gin.Context, params v1.ListPrivateNetworksParams) {
	networks, err := s.Engine.ListPrivateNetworks(params)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, networks)
}

// ReadPrivateNetwork ローカルネットワーク 詳細
// (GET /private_networks/{private_network_id}/)
func (s *Server) ReadPrivateNetwork(c *gin.Context, privateNetworkId v1.PrivateNetworkId) {
	network, err := s.Engine.ReadPrivateNetwork(privateNetworkId)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.ResponseBodyPrivateNetwork{
		PrivateNetwork: *network,
	})
}
