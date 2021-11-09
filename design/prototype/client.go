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

package prototype

import "github.com/sacloud/libphy/openapi"

const defaultAPIRootURL = "https://secure.sakura.ad.jp/cloud/api/dedicated-phy/1.0"

type Client struct {
	APIRootURL string

	AccessToken       string
	AccessTokenSecret string

	// TODO その他オプション類をここに追加
}

func (c *Client) serverURL() string {
	if c.APIRootURL != "" {
		return c.APIRootURL
	}
	return defaultAPIRootURL
}

func (c *Client) apiClient() (openapi.ClientWithResponsesInterface, error) {
	return openapi.NewClientWithResponses(
		c.serverURL(),
		func(client *openapi.Client) error {
			client.RequestEditors = []openapi.RequestEditorFn{
				openapi.PhyAuthInterceptor(c.AccessToken, c.AccessTokenSecret),
				openapi.PhyRequestInterceptor(),
			}
			return nil
		},
	)
}
