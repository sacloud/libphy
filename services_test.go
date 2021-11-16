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

package phy

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	v1 "github.com/sacloud/phy-go/apis/v1"
	"github.com/sacloud/phy-go/pointer"
	"github.com/sacloud/phy-go/stub"
	"github.com/stretchr/testify/require"
)

func TestServiceOp_ListServices(t *testing.T) {
	onlyUnitTest(t)

	tests := []struct {
		name    string
		args    *v1.ListServicesParams
		want    *v1.Services
		wantErr bool
	}{
		{
			name: "minimum",
			args: &v1.ListServicesParams{},
			want: &v1.Services{
				Meta: v1.PaginateMeta{
					Count: 1,
				},
				Services: []v1.Service{
					*testValueService01,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServiceOp{
				client: testClient(t),
			}
			got, err := op.List(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListServices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestServiceOp_ReadService(t *testing.T) {
	onlyUnitTest(t)

	tests := []struct {
		name      string
		serviceId v1.ServiceId
		want      *v1.Service
		wantErr   bool
	}{
		{
			name:      "minimum",
			serviceId: v1.ServiceId(testValueService01.ServiceId),
			want:      testValueService01,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServiceOp{
				client: testClient(t),
			}
			got, err := op.Read(context.Background(), tt.serviceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestServiceOp_UpdateService(t *testing.T) {
	onlyUnitTest(t)

	type args struct {
		serviceId v1.ServiceId
		params    v1.UpdateServiceParameter
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.Service
		wantErr bool
	}{
		{
			name: "minimum",
			args: args{
				serviceId: v1.ServiceId(testValueService01.ServiceId),
				params: v1.UpdateServiceParameter{
					Description: pointer.String("description01-upd"),
					Nickname:    "service01-upd",
				},
			},
			want: &v1.Service{
				Activated:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				Description: pointer.String("description01-upd"),
				Nickname:    "service01-upd",
				OptionPlans: nil,
				Plan: &v1.ServicePlan{
					Name:   "plan01",
					PlanId: "maker-series-spec-region-01",
				},
				ProductCategory: v1.ServiceProductCategoryServer,
				ServiceId:       "100000000001",
				Tags: []v1.Tag{
					{
						Color: pointer.String("tag1"),
						Label: "label",
						TagId: 1,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &ServiceOp{
				client: testClient(t),
			}
			got, err := op.Update(context.Background(), tt.args.serviceId, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestService_ErrorHandling(t *testing.T) {
	onlyUnitTest(t)

	stubServer := &stub.Server{
		ReadServiceFunc: func(c *gin.Context, serviceId v1.ServiceId) {
			c.JSON(http.StatusNotFound, &v1.ProblemDetails404{
				Detail: "not found detail",
				Status: http.StatusNotFound,
				Title:  v1.ProblemDetails404TitleNotFound,
				Type:   "about:blank",
			})
		},
	}
	httpServer := httptest.NewServer(stubServer.Handler())
	client := testClient(t)
	client.APIRootURL = httpServer.URL

	serviceOp := NewServiceOp(client)

	got, err := serviceOp.Read(context.Background(), v1.ServiceId("100000000001"))
	require.Nil(t, got)
	require.Error(t, err)
}
