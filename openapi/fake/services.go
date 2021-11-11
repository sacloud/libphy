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

// GetServices サービス一覧
// (GET /services/)
func (s *Server) GetServices(c *gin.Context, params openapi.GetServicesParams) {

}

// GetServicesServiceId サービス 詳細
// (GET /services/{service_id}/)
func (s *Server) GetServicesServiceId(c *gin.Context, serviceId openapi.ServiceId) {

}

// PatchServicesServiceId サービスの名称・説明の変更
// (PATCH /services/{service_id}/)
func (s *Server) PatchServicesServiceId(c *gin.Context, serviceId openapi.ServiceId, params openapi.PatchServicesServiceIdParams) {

}
