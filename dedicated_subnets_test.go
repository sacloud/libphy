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

package phy

import (
	"context"
	"testing"

	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/stretchr/testify/require"
)

func TestDedicatedSubnetOp_List(t *testing.T) {
	onlyUnitTest(t)

	subnets := testServer.Engine.GetDedicatedSubnets()
	tests := []struct {
		name    string
		params  *v1.ListDedicatedSubnetsParams
		want    *v1.DedicatedSubnets
		wantErr bool
	}{
		{
			name:   "minimum",
			params: &v1.ListDedicatedSubnetsParams{},
			want: &v1.DedicatedSubnets{
				Meta: v1.PaginateMeta{
					Count: len(subnets),
				},
				DedicatedSubnets: []v1.DedicatedSubnet{
					*subnets[0],
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &DedicatedSubnetOp{
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

func TestDedicatedSubnetOp_Read(t *testing.T) {
	onlyUnitTest(t)

	ds := testServer.Engine.GetDedicatedSubnets()[0]
	type args struct {
		dedicatedSubnetId v1.DedicatedSubnetId
		refresh           bool
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.DedicatedSubnet
		wantErr bool
	}{
		{
			name: "minimum",
			args: args{
				dedicatedSubnetId: ds.DedicatedSubnetId,
				refresh:           false,
			},
			want:    ds,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &DedicatedSubnetOp{
				client: testClient(t),
			}
			got, err := op.Read(context.Background(), tt.args.dedicatedSubnetId, tt.args.refresh)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.EqualValues(t, tt.want, got)
		})
	}
}
