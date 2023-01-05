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

	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

// ServiceAPI サービス関連API
type ServiceAPI interface {
	List(ctx context.Context, params *v1.ListServicesParams) (*v1.Services, error)
	Read(ctx context.Context, serviceId v1.ServiceId) (*v1.Service, error)
	Update(ctx context.Context, serviceId v1.ServiceId, params v1.UpdateServiceParameter) (*v1.Service, error)
}

// ServiceOp サービス関連の操作
type ServiceOp struct {
	client *Client
}

func NewServiceOp(client *Client) ServiceAPI {
	return &ServiceOp{client: client}
}

func (op *ServiceOp) List(ctx context.Context, params *v1.ListServicesParams) (*v1.Services, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	response, err := apiClient.ListServicesWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return response.Result()
}

func (op *ServiceOp) Read(ctx context.Context, serviceId v1.ServiceId) (*v1.Service, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	response, err := apiClient.ReadServiceWithResponse(ctx, serviceId)
	if err != nil {
		return nil, err
	}
	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.Service, nil
}

func (op *ServiceOp) Update(ctx context.Context, serviceId v1.ServiceId, params v1.UpdateServiceParameter) (*v1.Service, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	headers := &v1.UpdateServiceParams{
		XRequestedWith: v1.UpdateServiceParamsXRequestedWith(v1.XMLHttpRequest),
	}
	response, err := apiClient.UpdateServiceWithResponse(ctx, serviceId, headers, params)
	if err != nil {
		return nil, err
	}

	result, err := response.Result()
	if err != nil {
		return nil, err
	}
	return &result.Service, nil
}
