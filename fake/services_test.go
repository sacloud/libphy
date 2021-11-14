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
	"testing"
	"time"

	"github.com/sacloud/phy-go/openapi"
	"github.com/sacloud/phy-go/pointer"
	"github.com/stretchr/testify/require"
)

func TestDataStore_Services(t *testing.T) {
	ds := &Engine{
		Services: []*openapi.Service{
			{
				Activated:   time.Now(),
				Description: pointer.String("description1"),
				Nickname:    "nickname1",
				OptionPlans: nil,
				Plan: &openapi.ServicePlan{
					Name:   "plan-01",
					PlanId: "maker-series-spec-region-01",
				},
				ProductCategory: openapi.ServiceProductCategoryServer,
				ServiceId:       "100000000001",
				Tags: []openapi.Tag{
					{
						Color: pointer.String("ffffff"),
						Label: "label",
						TagId: 1,
					},
				},
			},
			{
				Activated:   time.Now(),
				Description: pointer.String("description2"),
				Nickname:    "nickname2",
				OptionPlans: nil,
				Plan: &openapi.ServicePlan{
					Name:   "plan-02",
					PlanId: "maker-series-spec-region-02",
				},
				ProductCategory: openapi.ServiceProductCategoryServer,
				ServiceId:       "100000000002",
				Tags: []openapi.Tag{
					{
						Color: pointer.String("ffffff"),
						Label: "label",
						TagId: 1,
					},
				},
			},
		},
	}

	t.Run("select all", func(t *testing.T) {
		services, err := ds.ListServices(openapi.ListServicesParams{})
		require.NoError(t, err)
		require.NotNil(t, services)
		require.Equal(t, services.Meta.Count, len(ds.Services))
		require.Len(t, services.Services, len(ds.Services))
	})

	t.Run("select by id", func(t *testing.T) {
		service, err := ds.ReadService("100000000001")
		require.NoError(t, err)
		require.NotNil(t, service)
		require.Equal(t, "nickname1", service.Nickname)
	})

	t.Run("update", func(t *testing.T) {
		svc, err := ds.UpdateService("100000000001", openapi.UpdateServiceJSONBody{
			Description: nil,
			Nickname:    pointer.String("nickname1-upd"),
		})
		require.NoError(t, err)
		require.NotNil(t, svc)

		upd, err := ds.ReadService("100000000001")
		require.NoError(t, err)

		require.Equal(t, "nickname1-upd", upd.Nickname)
	})
}
