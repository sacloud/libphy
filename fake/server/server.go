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

package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/sacloud/phy-go/fake"
	"github.com/sacloud/phy-go/openapi"
)

// Server PHY APIのFakeサーバ実装
//
// インメモリデータソースを備え、正常系のシナリオを中心に実装。
// エラー処理が簡易に実装されており実サーバとは返すエラーが違うことがあるため
//  異常系のテストをしたい場合は代わりにstubパッケージを利用してください。
type Server struct {
	Engine *fake.Engine

	httpServer *httptest.Server
}

// Start テスト用のHTTPサーバを起動し、クリーンアップ用のfuncを返す
func (s *Server) Start() func() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	s.httpServer = httptest.NewServer(openapi.RegisterHandlers(router, s))
	return s.httpServer.Close
}

// URL テスト用サーバのURLを返す
func (s *Server) URL() string {
	if s.httpServer != nil {
		return s.httpServer.URL
	}
	panic("http server is not started")
}

func (s *Server) handleError(c *gin.Context, err error) {
	if c == nil || err == nil {
		panic("invalid arguments")
	}

	if engineErr, ok := err.(*fake.Error); ok {
		switch engineErr.Type {
		case fake.ErrorTypeInvalidRequest:
			c.JSON(http.StatusBadRequest, &openapi.ProblemDetails400{
				Detail: engineErr.Error(),
				Status: http.StatusBadRequest,
				Title:  openapi.ProblemDetails400TitleInvalid, // この実装ではinvalid固定
				Type:   "about:blank",
				InvalidParameters: &openapi.InvalidParameter{
					NonFieldErrors: &openapi.InvalidParameterDetails{
						{
							Code:    "xxx",
							Message: engineErr.Error(),
						},
					},
				},
			})
		case fake.ErrorTypeNotFound:
			c.JSON(http.StatusNotFound, &openapi.ProblemDetails404{
				Detail: engineErr.Error(),
				Status: http.StatusNotFound,
				Title:  openapi.ProblemDetails404TitleNotFound,
				Type:   "about:blank",
			})
		case fake.ErrorTypeConflict:
			c.JSON(http.StatusConflict, &openapi.ProblemDetails409{
				Detail: engineErr.Error(),
				Status: http.StatusConflict,
				Title:  openapi.ProblemDetails409TitleConflict,
				Type:   "about:blank",
			})
		}
	}

	// Note: API定義的には503が返ることもあるがこの実装では5xx系は全てInternalServerErrorとして扱う
	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("unknown error: %s", err)})
}
