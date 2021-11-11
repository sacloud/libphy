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

// GetDedicatedSubnets 専用グローバルネットワーク 一覧
// (GET /dedicated_subnets/)
func (s *Server) GetDedicatedSubnets(c *gin.Context, params openapi.GetDedicatedSubnetsParams) {

}

// GetDedicatedSubnetsDedicatedSubnetId 専用グローバルネットワーク
// (GET /dedicated_subnets/{dedicated_subnet_id}/)
func (s *Server) GetDedicatedSubnetsDedicatedSubnetId(c *gin.Context, dedicatedSubnetId openapi.DedicatedSubnetId, params openapi.GetDedicatedSubnetsDedicatedSubnetIdParams) {

}
