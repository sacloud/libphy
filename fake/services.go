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
	v1 "github.com/sacloud/phy-go/apis/v1"
)

// ListServices サービス一覧
// (GET /services/)
func (engine *Engine) ListServices(_ v1.ListServicesParams) (*v1.Services, error) {
	defer engine.rLock()()

	// TODO 検索条件の処理を実装

	return &v1.Services{
		Meta: v1.PaginateMeta{
			Count: len(engine.Services),
		},
		Services: engine.services(),
	}, nil
}

// ReadService サービス 詳細
// (GET /services/{service_id}/)
func (engine *Engine) ReadService(serviceId v1.ServiceId) (*v1.Service, error) {
	defer engine.rLock()()
	s := engine.getServiceById(serviceId)
	if s != nil {
		// パッケージ外に返す時はディープコピーしたものを返す
		var service v1.Service
		if err := deepcopy.Copy(&service, s); err != nil {
			return nil, err
		}
		return &service, nil
	}
	return nil, fmt.Errorf("service %q not found", serviceId)
}

// UpdateService サービスの名称・説明の変更
// (PATCH /services/{service_id}/)
func (engine *Engine) UpdateService(serviceId v1.ServiceId, body v1.UpdateServiceParameter) (*v1.Service, error) {
	defer engine.lock()()
	service := engine.getServiceById(serviceId)
	if service != nil {
		service.Nickname = body.Nickname
		service.Description = body.Description
		var svc v1.Service
		if err := deepcopy.Copy(&svc, service); err != nil {
			return nil, err
		}
		return &svc, nil
	}
	return nil, fmt.Errorf("service %q not found", serviceId)
}

// services []*v1.Serviceから[]v1.Serviceに変換して返す
func (engine *Engine) services() []v1.Service {
	var results []v1.Service
	for _, s := range engine.Services {
		results = append(results, *s)
	}
	return results
}

func (engine *Engine) getServiceById(serviceId v1.ServiceId) *v1.Service {
	for _, s := range engine.Services {
		if s.ServiceId == string(serviceId) {
			return s
		}
	}
	return nil
}
