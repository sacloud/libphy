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

package phy_test

import (
	"context"
	"fmt"

	"github.com/sacloud/phy-api-go"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

var serverURL = phy.DefaultAPIRootURL

func Example() {
	client := &phy.Client{
		APIRootURL: serverURL, // 省略可
	}

	// サーバ操作
	serverOp := phy.NewServerOp(client)
	found, err := serverOp.List(context.Background(), &v1.ListServersParams{})
	if err != nil {
		panic(err)
	}
	for _, sv := range found.Servers {
		fmt.Println(sv.Service.Nickname)
	}

	// output:
	// server01
}

func ExampleServerAPI() {
	client := &phy.Client{
		APIRootURL: serverURL, // 省略可
	}

	// サーバ操作
	ctx := context.Background()
	serverOp := phy.NewServerOp(client)
	found, err := serverOp.List(ctx, &v1.ListServersParams{})
	if err != nil {
		panic(err)
	}

	// 電源がONのサーバをシャットダウン
	for _, sv := range found.Servers {
		if sv.CachedPowerStatus.Status == v1.CachedPowerStatusStatusOn {
			// v1.ServerPowerOperationsSoft == ACPIシャットダウン
			if err := serverOp.PowerControl(ctx, sv.ServerId, v1.ServerPowerOperationsSoft); err != nil {
				panic(err)
			}
			fmt.Printf("shutting down: %s\n", sv.Service.Nickname)
		}
	}

	// output:
	// shutting down: server01
}

func ExampleServiceAPI() {
	client := &phy.Client{
		APIRootURL: serverURL, // 省略可
	}

	// サービス一覧取得
	serviceOp := phy.NewServiceOp(client)
	found, err := serviceOp.List(context.Background(), &v1.ListServicesParams{})
	if err != nil {
		panic(err)
	}
	for _, svc := range found.Services {
		fmt.Println(svc.Nickname)
	}

	// output:
	// server01
}
