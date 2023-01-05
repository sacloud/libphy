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

package fake

import (
	"github.com/getlantern/deepcopy"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

// ListDedicatedSubnets 専用グローバルネットワーク 一覧
// (GET /dedicated_subnets/)
func (engine *Engine) ListDedicatedSubnets(params v1.ListDedicatedSubnetsParams) (*v1.DedicatedSubnets, error) {
	defer engine.rLock()()

	// TODO 検索条件の処理を実装

	return &v1.DedicatedSubnets{
		Meta: v1.PaginateMeta{
			Count: len(engine.DedicatedSubnets),
		},
		DedicatedSubnets: engine.dedicatedSubnets(),
	}, nil
}

// ReadDedicatedSubnet 専用グローバルネットワーク
// (GET /dedicated_subnets/{dedicated_subnet_id}/)
// Note: paramの処理は未実装
func (engine *Engine) ReadDedicatedSubnet(dedicatedSubnetId v1.DedicatedSubnetId, _ v1.ReadDedicatedSubnetParams) (*v1.DedicatedSubnet, error) {
	defer engine.rLock()()

	d := engine.getDedicatedSubnetById(dedicatedSubnetId)
	if d != nil {
		// パッケージ外に返す時はディープコピーしたものを返す
		var subnet v1.DedicatedSubnet
		if err := deepcopy.Copy(&subnet, d); err != nil {
			return nil, err
		}
		return &subnet, nil
	}
	return nil, NewError(ErrorTypeNotFound, "dedicated-subnet", dedicatedSubnetId)
}

// dedicatedSubnets []*v1.DedicatedSubnetから[]v1.DedicatedSubnetに変換して返す
func (engine *Engine) dedicatedSubnets() []v1.DedicatedSubnet {
	var results []v1.DedicatedSubnet
	for _, d := range engine.DedicatedSubnets {
		results = append(results, *d)
	}
	return results
}

func (engine *Engine) getDedicatedSubnetById(dedicatedSubnetId v1.DedicatedSubnetId) *v1.DedicatedSubnet {
	for _, d := range engine.DedicatedSubnets {
		if d.DedicatedSubnetId == dedicatedSubnetId {
			return d
		}
	}
	return nil
}
