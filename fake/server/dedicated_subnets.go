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

// ListDedicatedSubnets 専用グローバルネットワーク 一覧
// (GET /dedicated_subnets/)
func (s *Server) ListDedicatedSubnets(c *gin.Context, params v1.ListDedicatedSubnetsParams) {
	subnets, err := s.Engine.ListDedicatedSubnets(params)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, subnets)
}

// ReadDedicatedSubnet 専用グローバルネットワーク
// (GET /dedicated_subnets/{dedicated_subnet_id}/)
func (s *Server) ReadDedicatedSubnet(c *gin.Context, dedicatedSubnetId v1.DedicatedSubnetId, params v1.ReadDedicatedSubnetParams) {
	subnet, err := s.Engine.ReadDedicatedSubnet(dedicatedSubnetId, params)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.ResponseBodyDedicatedSubnet{
		DedicatedSubnet: *subnet,
	})
}
