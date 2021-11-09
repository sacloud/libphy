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
	"os"
	"testing"
)

func TestDummy(t *testing.T) {
	if os.Getenv("TESTACC") == "" {
		t.Skip("required: environment variable 'TESTACC'")
	}

	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")
	if token == "" || secret == "" {
		t.Skip("required: environment variable 'SAKURACLOUD_ACCESS_TOKEN' and 'SAKURACLOUD_ACCESS_TOKEN_SECRET'")
	}

	client, err := NewClientWithResponses("https://secure.sakura.ad.jp/cloud/api/dedicated-phy/1.0", func(c *Client) error {
		c.RequestEditors = []RequestEditorFn{
			func(ctx context.Context, req *http.Request) error {
				req.SetBasicAuth(os.Getenv("SAKURACLOUD_ACCESS_TOKEN"), os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET"))
				req.Header.Add("X-Requested-With", "XMLHttpRequest")
				return nil
			},
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	servers, err := client.GetServersWithResponse(context.Background(), &GetServersParams{})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(servers.Body))
}
