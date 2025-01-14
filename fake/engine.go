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

package fake

import (
	"encoding/json"
	"sync"
	"time"

	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

const defaultActionInterval = 100 * time.Millisecond

// Engine Fakeサーバであつかうダミーデータを表す
//
// Serverに渡した後は各フィールドを外部から操作しないこと
// 各フィールドの値を参照したい場合はGetxxx()を用いること
type Engine struct {
	Services         []*v1.Service
	Servers          []*Server
	DedicatedSubnets []*v1.DedicatedSubnet
	PrivateNetworks  []*v1.PrivateNetwork

	// ActionInterval バックグラウンドでリソースの状態を変化させるアクションの実行間隔
	ActionInterval time.Duration

	// GeneratedID 採番済みの最終ID
	//
	// DataStoreの各フィールドの値との整合性は確認されないため利用者側が管理する必要がある
	GeneratedID int

	mu sync.RWMutex
}

func (engine *Engine) GetServices() []*v1.Service {
	defer engine.rLock()()
	var results []*v1.Service
	data, err := json.Marshal(engine.Services)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &results); err != nil {
		panic(err)
	}
	return results
}

func (engine *Engine) GetServers() []*Server {
	defer engine.rLock()()
	var results []*Server
	data, err := json.Marshal(engine.Servers)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &results); err != nil {
		panic(err)
	}
	return results
}

func (engine *Engine) GetDedicatedSubnets() []*v1.DedicatedSubnet {
	defer engine.rLock()()
	var results []*v1.DedicatedSubnet
	data, err := json.Marshal(engine.DedicatedSubnets)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &results); err != nil {
		panic(err)
	}
	return results
}

func (engine *Engine) GetPrivateNetworks() []*v1.PrivateNetwork {
	defer engine.rLock()()
	var results []*v1.PrivateNetwork
	data, err := json.Marshal(engine.PrivateNetworks)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &results); err != nil {
		panic(err)
	}
	return results
}

func (engine *Engine) lock() func() {
	engine.mu.Lock()
	return engine.mu.Unlock
}

func (engine *Engine) rLock() func() {
	engine.mu.RLock()
	return engine.mu.RUnlock
}

// nextId GeneratedIDを+1したものを返す
//
// ロックは行わないため呼び出し側で適切に制御すること
func (engine *Engine) nextId() int {
	engine.GeneratedID++
	id := engine.GeneratedID
	return id
}

func (engine *Engine) actionInterval() time.Duration {
	if engine.ActionInterval > 0 {
		return engine.ActionInterval
	}
	return defaultActionInterval
}

func (engine *Engine) startUpdateAction(action func()) {
	time.Sleep(engine.actionInterval())
	defer engine.lock()()
	action()
}
