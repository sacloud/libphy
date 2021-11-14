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
	"fmt"

	"github.com/getlantern/deepcopy"
	"github.com/sacloud/phy-go/openapi"
)

// ListServices サービス一覧
// (GET /services/)
func (engine *Engine) ListServices(_ openapi.ListServicesParams) (*openapi.Services, error) {
	defer engine.rLock()()

	// TODO 検索条件の処理を実装

	return &openapi.Services{
		Meta: openapi.PaginateMeta{
			Count: len(engine.Services),
		},
		Services: engine.services(),
	}, nil
}

// ReadService サービス 詳細
// (GET /services/{service_id}/)
func (engine *Engine) ReadService(serviceId openapi.ServiceId) (*openapi.Service, error) {
	defer engine.rLock()()
	s := engine.getServiceById(serviceId)
	if s != nil {
		// パッケージ外に返す時はディープコピーしたものを返す
		var service openapi.Service
		if err := deepcopy.Copy(&service, s); err != nil {
			return nil, err
		}
		return &service, nil
	}
	return nil, fmt.Errorf("service %q not found", serviceId)
}

// UpdateService サービスの名称・説明の変更
// (PATCH /services/{service_id}/)
func (engine *Engine) UpdateService(serviceId openapi.ServiceId, body openapi.UpdateServiceJSONBody) (*openapi.Service, error) {
	defer engine.lock()()
	service := engine.getServiceById(serviceId)
	if service != nil {
		if body.Nickname != nil {
			service.Nickname = *body.Nickname
		}
		if body.Description != nil {
			service.Description = body.Description
		}
		var svc openapi.Service
		if err := deepcopy.Copy(&svc, service); err != nil {
			return nil, err
		}
		return &svc, nil
	}
	return nil, fmt.Errorf("service %q not found", serviceId)
}

// services []*openapi.Serviceから[]openapi.Serviceに変換して返す
func (engine *Engine) services() []openapi.Service {
	var results []openapi.Service
	for _, s := range engine.Services {
		results = append(results, *s)
	}
	return results
}

func (engine *Engine) getServiceById(serviceId openapi.ServiceId) *openapi.Service {
	for _, s := range engine.Services {
		if s.ServiceId == string(serviceId) {
			return s
		}
	}
	return nil
}
