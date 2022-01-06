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

package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	v1 "github.com/sacloud/phy-go/apis/v1"
	"github.com/sacloud/phy-go/fake"
)

// Server PHY APIのFakeサーバ実装
//
// インメモリデータソースを備え、正常系のシナリオを中心に実装。
// エラー処理が簡易に実装されており実サーバとは返すエラーが違うことがあるため
//  異常系のテストをしたい場合は代わりにstubパッケージを利用してください。
type Server struct {
	Engine *fake.Engine
}

func (s *Server) Handler() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	if os.Getenv("PHY_SERVER_DEBUG") != "" {
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	if os.Getenv("PHY_SERVER_LOGGING") != "" {
		engine.Use(gin.Logger())
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return v1.RegisterHandlers(engine, s)
}

func (s *Server) handleError(c *gin.Context, err error) {
	if c == nil || err == nil {
		panic("invalid arguments")
	}

	if engineErr, ok := err.(*fake.Error); ok {
		switch engineErr.Type {
		case fake.ErrorTypeInvalidRequest:
			c.JSON(http.StatusBadRequest, &v1.ProblemDetails400{
				Detail: engineErr.Error(),
				Status: http.StatusBadRequest,
				Title:  v1.ProblemDetails400TitleInvalid, // この実装ではinvalid固定
				Type:   "about:blank",
				InvalidParameters: &v1.InvalidParameter{
					NonFieldErrors: &v1.InvalidParameterDetails{
						{
							Code:    "xxx",
							Message: engineErr.Error(),
						},
					},
				},
			})
		case fake.ErrorTypeNotFound:
			c.JSON(http.StatusNotFound, &v1.ProblemDetails404{
				Detail: engineErr.Error(),
				Status: http.StatusNotFound,
				Title:  v1.ProblemDetails404TitleNotFound,
				Type:   "about:blank",
			})
		case fake.ErrorTypeConflict:
			c.JSON(http.StatusConflict, &v1.ProblemDetails409{
				Detail: engineErr.Error(),
				Status: http.StatusConflict,
				Title:  v1.ProblemDetails409TitleConflict,
				Type:   "about:blank",
			})
		}
	}

	// Note: API定義的には503が返ることもあるがこの実装では5xx系は全てInternalServerErrorとして扱う
	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("unknown error: %s", err)})
}
