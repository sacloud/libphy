// Copyright 2021 The libphy authors
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

package openapi

import (
	"context"
	"net/http"
)

// PhyAuthInterceptor PHYへのリクエストに認証情報の注入を行う
func PhyAuthInterceptor(token, secret string) func(context.Context, *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth(token, secret)
		return nil
	}
}

// PhyRequestInterceptor PHYへのリクエストに必要なヘッダ類の注入を行う
func PhyRequestInterceptor() func(context.Context, *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		// https://manual.sakura.ad.jp/ds/phy/api/api-spec.html#section/基本的な使い方/入力パラメーター
		// > GET 以外のメソッドでは、CSRFの対策として、ヘッダーに X-Requested-With: XMLHttpRequest を指定します。
		if req.Method != http.MethodGet {
			req.Header.Add("X-Requested-With", "XMLHttpRequest")
		}
		return nil
	}
}
