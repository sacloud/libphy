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

package fake

import (
	"fmt"

	"github.com/getlantern/deepcopy"
	v1 "github.com/sacloud/phy-go/apis/v1"
)

// ListPrivateNetworks ローカルネットワーク 一覧
// (GET /private_networks/)
func (engine *Engine) ListPrivateNetworks(params v1.ListPrivateNetworksParams) (*v1.PrivateNetworks, error) {
	defer engine.rLock()()

	// TODO 検索条件の処理を実装

	return &v1.PrivateNetworks{
		Meta: v1.PaginateMeta{
			Count: len(engine.PrivateNetworks),
		},
		PrivateNetworks: engine.privateNetworks(),
	}, nil
}

// ReadPrivateNetwork ローカルネットワーク 詳細
// (GET /private_networks/{private_network_id}/)
func (engine *Engine) ReadPrivateNetwork(privateNetworkId v1.PrivateNetworkId) (*v1.PrivateNetwork, error) {
	defer engine.rLock()()

	pn := engine.getPrivateNetworkById(privateNetworkId)
	if pn != nil {
		// パッケージ外に返す時はディープコピーしたものを返す
		var network v1.PrivateNetwork
		if err := deepcopy.Copy(&network, pn); err != nil {
			return nil, err
		}
		return &network, nil
	}
	return nil, fmt.Errorf("private network %q not found", privateNetworkId)
}

// privateNetworks []*v1.PrivateNetworkから[]v1.PrivateNetworkに変換して返す
func (engine *Engine) privateNetworks() []v1.PrivateNetwork {
	var results []v1.PrivateNetwork
	for _, p := range engine.PrivateNetworks {
		results = append(results, *p)
	}
	return results
}

func (engine *Engine) getPrivateNetworkById(privateNetworkId v1.PrivateNetworkId) *v1.PrivateNetwork {
	for _, p := range engine.PrivateNetworks {
		if p.PrivateNetworkId == string(privateNetworkId) {
			return p
		}
	}
	return nil
}
