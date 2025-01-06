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

package phy

import (
	"context"
	"testing"

	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/stretchr/testify/require"
)

func TestPrivateNetworkOp_List(t *testing.T) {
	onlyUnitTest(t)
	networks := testServer.Engine.GetPrivateNetworks()

	tests := []struct {
		name    string
		params  *v1.ListPrivateNetworksParams
		want    *v1.PrivateNetworks
		wantErr bool
	}{
		{
			name:   "minimum",
			params: &v1.ListPrivateNetworksParams{},
			want: &v1.PrivateNetworks{
				Meta: v1.PaginateMeta{
					Count: len(networks),
				},
				PrivateNetworks: []v1.PrivateNetwork{
					*networks[0],
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &PrivateNetworkOp{
				client: testClient(t),
			}
			got, err := op.List(context.Background(), tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestPrivateNetworkOp_Read(t *testing.T) {
	onlyUnitTest(t)
	networks := testServer.Engine.GetPrivateNetworks()

	tests := []struct {
		name             string
		privateNetworkId v1.PrivateNetworkId
		want             *v1.PrivateNetwork
		wantErr          bool
	}{
		{
			name:             "minimum",
			privateNetworkId: networks[0].PrivateNetworkId,
			want:             networks[0],
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &PrivateNetworkOp{
				client: testClient(t),
			}
			got, err := op.Read(context.Background(), tt.privateNetworkId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}
