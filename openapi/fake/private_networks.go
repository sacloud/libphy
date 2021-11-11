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

// GetPrivateNetworks ローカルネットワーク 一覧
// (GET /private_networks/)
func (s *Server) GetPrivateNetworks(c *gin.Context, params openapi.GetPrivateNetworksParams) {

}

// GetPrivateNetworksPrivateNetworkId ローカルネットワーク 詳細
// (GET /private_networks/{private_network_id}/)
func (s *Server) GetPrivateNetworksPrivateNetworkId(c *gin.Context, privateNetworkId openapi.PrivateNetworkId) {

}
