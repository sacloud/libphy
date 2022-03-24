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

package phy

import (
	"context"

	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

// DedicatedSubnetAPI 専用グローバルネットワーク関連API
type DedicatedSubnetAPI interface {
	List(ctx context.Context, params *v1.ListDedicatedSubnetsParams) (*v1.DedicatedSubnets, error)
	Read(ctx context.Context, dedicatedSubnetId v1.DedicatedSubnetId, refresh bool) (*v1.DedicatedSubnet, error)
}

type DedicatedSubnetOp struct {
	client *Client
}

func NewDedicatedSubnetOp(client *Client) DedicatedSubnetAPI {
	return &DedicatedSubnetOp{client: client}
}

func (op *DedicatedSubnetOp) List(ctx context.Context, params *v1.ListDedicatedSubnetsParams) (*v1.DedicatedSubnets, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	response, err := apiClient.ListDedicatedSubnetsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return response.Result()
}

func (op *DedicatedSubnetOp) Read(ctx context.Context, dedicatedSubnetId v1.DedicatedSubnetId, refresh bool) (*v1.DedicatedSubnet, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}

	params := &v1.ReadDedicatedSubnetParams{Refresh: &refresh}
	response, err := apiClient.ReadDedicatedSubnetWithResponse(ctx, dedicatedSubnetId, params)
	if err != nil {
		return nil, err
	}
	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.DedicatedSubnet, nil
}
