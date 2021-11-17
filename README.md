# phy-go 

[![Go Reference](https://pkg.go.dev/badge/github.com/sacloud/phy-go.svg)](https://pkg.go.dev/github.com/sacloud/phy-go)
[![Tests](https://github.com/sacloud/phy-go/workflows/Tests/badge.svg)](https://github.com/sacloud/phy-go/actions/workflows/tests.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sacloud/phy-go)](https://goreportcard.com/report/github.com/sacloud/phy-go)

Library for PHY API / [さくらの専用サーバPHY](https://server.sakura.ad.jp) の [API](https://manual.sakura.ad.jp/ds/phy/api/api-spec.html) をGo言語から扱うためのライブラリ

PHY API: [https://manual.sakura.ad.jp/ds/phy/api/api-spec.html](https://manual.sakura.ad.jp/ds/phy/api/api-spec.html)

## Overview

[oapi-codegen](https://github.com/deepmap/oapi-codegen) によって生成されたGoのコードに加え、
Fake/Stubサーバの実装やより簡易に使えるようにラップしたクライアントコードを提供します。

:warning: phy-goは現在開発中です。v1に達するまでは後方互換性のない変更が行われる可能性があります。

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

## Fakeサーバ(スタンドアロン)

FakeサーバはGoのコード以外からでも利用できるようにスタンドアロン版を提供しています。

### Fakeサーバのインストール

リリースページからダウンロード or `go install`してください。

リリースページ: [https://github.com/sacloud/phy-go/releases/latest](https://github.com/sacloud/phy-go/releases/latest)

```bash
go install github.com/sacloud/phy-go/cmd/phy-go-fake-server
```

### Fakeサーバの利用方法

```bash
$ phy-go-fake-server --help
Start the web server

Usage:
  phy-go-fake-server [flags]

Flags:
      --addr string      the address for the server to listen on (default ":8080")
      --data string      the file path to the fake data JSON file
  -h, --help             help for phy-go-fake-server
      --output-example   the flag to output a fake data JSON example
  -v, --version          version for phy-go-fake-server
```

- `--addr`: Fakeサーバがリッスンするアドレス
- `--data`: FakeデータのJSONファイルへのパス、省略した場合はデフォルトのダミーデータが利用される
- `--output-example`: FakeデータのJSONファイルの例を出力

起動したら次のようにリクエストを行えます。

```bash
# localhost:8080で起動した場合の例
$ curl http://localhost:8080/services/
```

### Fakeデータのカスタマイズ

`--output-example`でJSONファイルの雛形を出力し、編集、その後`--data`でファイルパスを指定します。

```bash
# 雛形を出力
$ phy-go-fake-server --output-example > fake.json
# 編集
$ vi fake.json
# データファイルのパスを指定して起動
$ phy-go-fake-server --data=fake.json
```

## License

`phy-go` Copyright 2021 The phy-go authors.

This project is published under [Apache 2.0 License](LICENSE).
