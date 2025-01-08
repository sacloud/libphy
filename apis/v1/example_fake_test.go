// Copyright 2021-2025 The phy-api-go authors
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

package v1_test

import (
	"net/http/httptest"
	"time"

	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/sacloud/phy-api-go/fake"
	"github.com/sacloud/phy-api-go/fake/server"
)

func init() {
	initFakeServer()
}

func initFakeServer() {
	fakeServer := &server.Server{
		Engine: &fake.Engine{
			Services: []*v1.Service{
				{
					Activated: time.Now(),
					Nickname:  "server01",
					Plan: &v1.ServicePlan{
						Name:   "plan01",
						PlanId: "xxx",
					},
					ProductCategory: v1.ServiceProductCategoryServer,
					ServiceId:       "100000000001",
				},
			},
		},
	}
	sv := httptest.NewServer(fakeServer.Handler())
	serverURL = sv.URL
}
