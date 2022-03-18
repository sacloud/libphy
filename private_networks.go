// Copyright 2021-2022 The phy-go authors
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

	v1 "github.com/sacloud/phy-go/apis/v1"
)

// PrivateNetworkAPI ローカルネットワーク関連API
type PrivateNetworkAPI interface {
	List(ctx context.Context, params *v1.ListPrivateNetworksParams) (*v1.PrivateNetworks, error)
	Read(ctx context.Context, privateNetworkId v1.PrivateNetworkId) (*v1.PrivateNetwork, error)
}

type PrivateNetworkOp struct {
	client *Client
}

func NewPrivateNetworkOp(client *Client) PrivateNetworkAPI {
	return &PrivateNetworkOp{client: client}
}

func (op *PrivateNetworkOp) List(ctx context.Context, params *v1.ListPrivateNetworksParams) (*v1.PrivateNetworks, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	response, err := apiClient.ListPrivateNetworksWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return response.Result()
}

func (op *PrivateNetworkOp) Read(ctx context.Context, privateNetworkId v1.PrivateNetworkId) (*v1.PrivateNetwork, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	response, err := apiClient.ReadPrivateNetworkWithResponse(ctx, privateNetworkId)
	if err != nil {
		return nil, err
	}
	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.PrivateNetwork, nil
}
