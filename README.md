# phy-go 

[![Go Reference](https://pkg.go.dev/badge/github.com/sacloud/phy-go.svg)](https://pkg.go.dev/github.com/sacloud/phy-go)
[![Tests](https://github.com/sacloud/phy-go/workflows/Tests/badge.svg)](https://github.com/sacloud/phy-go/actions/workflows/tests.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sacloud/phy-go)](https://goreportcard.com/report/github.com/sacloud/phy-go)

Library for PHY API / [さくらの専用サーバPHY](https://server.sakura.ad.jp) の [API](https://manual.sakura.ad.jp/ds/phy/api/api-spec.html) をGo言語から扱うためのライブラリ

PHY API: [https://manual.sakura.ad.jp/ds/phy/api/api-spec.html](https://manual.sakura.ad.jp/ds/phy/api/api-spec.html)

## Overview

[oapi-codegen](https://github.com/deepmap/oapi-codegen) によって生成されたGoのコードに加え、
Fake/Stubサーバの実装やより簡易に使えるようにラップしたクライアントコードを提供します。

#### phy-goを利用したクライアントコードの例

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sacloud/phy-go"
	v1 "github.com/sacloud/phy-go/apis/v1"
)

func main() {
	// APIキー
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	// httpクライアント
	client := &phy.Client{
		Token:  token,
		Secret: secret,
	}

	// サーバ一覧
	serverOp := phy.NewServerOp(client)
	found, err := serverOp.List(context.Background(), &v1.ListServersParams{})
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range found.Servers {
		fmt.Println(server.Service.Nickname)
	}
}
```

## Installation

Use go get.

    go get github.com/sacloud/phy-go

Then import the `phy` package into your own code.

    import "github.com/sacloud/phy-go"

## License

`phy-go` Copyright 2021 The phy-go authors.

This project is published under [Apache 2.0 License](LICENSE).