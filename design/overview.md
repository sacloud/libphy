# phy-go

## 概要

[さくらの専用サーバ PHY](https://server.sakura.ad.jp)が提供するAPIをGoから利用するためのライブラリを提供する。

## 開発の背景

PHYのAPI定義はOpenAPI 3.0仕様により文書化されている。

API定義: [https://manual.sakura.ad.jp/ds/phy/api/api-spec.html](https://manual.sakura.ad.jp/ds/phy/api/api-spec.html)

これを元にGo向けのコード生成を行うツールはいくつかあり、phy-goでは [oapi-codegen](https://github.com/deepmap/oapi-codegen) を利用している。

生成されたコードはそのまま利用可能ではあるが、トランスポートレベルの詳細を含んでいたり、高いカスタマイズ性の確保のためにやや冗長な記述が必要となっており、そのままではアプリケーションに組み込みづらい。

例: `oapi-codegen`で生成されたコードを用いてサーバの電源を操作する

```go
	result, err := client.PostServersServerIdPowerControlWithResponse(
		context.Background(),
		serverId,
		&PostServersServerIdPowerControlParams{
			XRequestedWith: "XMLHttpRequest",
		},
		PostServersServerIdPowerControlJSONRequestBody{
			Operation: "soft", // ACPIシャットダウン
		})
	if err != nil {
		return err
	}
	if result.StatusCode() == http.StatusNoContent {
		// ...
	}
```

これを以下のように扱いたい。

```go
	if err := serverAPI.Power(context.Background(), serverId, powerOps.Soft); err != nil {
		return err
	}
```

## 仕様/制約など

公開されているAPI定義の記述誤りやコード生成ツールの制約があり、API定義を手作業で修正する必要がある。  
また、API定義が網羅的に書かれているわけではなく、例えば値のバリデーションルールについて厳密に記載されていない。  

このためある程度のロジックをphy-go側で実装する必要がある。

## 設計/実装

### 実装方針

- API定義からコード生成ツールを用いてコード生成する。
- より簡潔に利用できるようにするために、ツールから生成されたコードをラップするコードを手作業で実装していく。
- 機械的に追加できるコードを実装するためにコード生成ツールのテンプレートを修正して利用する。
  ただしテンプレートの修正は煩雑になりがちなので、なるべくテンプレートの利用は控える。

### コード生成

コード生成には [https://github.com/deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen) を利用する。

ツール選定時には以下を検討した。

- [https://github.com/deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [https://github.com/OpenAPITools/openapi-generator](https://github.com/OpenAPITools/openapi-generator)

参考: https://github.com/sacloud/phy-go/issues/5

### パッケージ構造

- 実装するコードはプロジェクトルート直下の`phy`パッケージに
- 定義ファイル類はopenapiディレクトリ配下に
- 自動生成されるコードはopenapiディレクトリ配下に
- その他機能が増えるようであればプロジェクトルート配下にディレクトリを追加していく

```
.
├── xxx.go       # プロジェクト直下に各種実装コードを配置
├── go.mod
├── go.sum
└── openapi    # openapi配下に定義ファイルや自動生成されたコードを配置
    ├── spec
    │   └── swagger.yaml
    ├── test   # openapi用のFakeサーバの実装
    │   └── xxx.go
    ├── zz_client_gen.go
    └── zz_types_gen.go
```

#### `phy`パッケージ

- ツールにより生成されたコードをラップし、より簡潔に使えるようにしたもの
- 以下のようなコンポーネントを提供する
  - Client: APIキーなどのAPIを呼ぶにあたってのパラメータを保持し、HTTPクライアントとしてのインターフェースを提供する
  - xxxAPI: 各APIリソースを操作するためのインターフェイス
  - xxxOp: xxxAPIに対する実装、API定義で定義されたpathと1:1で対応する
    - ServerAPI / ServerOp
    - ServiceAPI / ServiceOp
    - GlobalNetworkAPI / GlobalNetworkOp
    - LocalNetworkAPI / LocalNetworkOp

#### `openapi`パッケージ

ツールにより生成されたコードを格納する。

### テスト

初期実装としてはlibsacloudのFakeドライバー(高レベルAPI単位でのテストダブル)を提供しない。  
代わりにOpenAPIレベルでのテストのためのFakeサーバー実装を提供する。

Fakeサーバはインメモリでデータを保持する簡易的な実装とする。

### sacloud/go-httpによるトランスポートレイヤでの機能の共通化

さくらのクラウド向けのAPIライブラリ [sacloud/libsacloud](https://github.com/sacloud/libsacloud) からトランスポートレイヤでの処理を切り出した [sacloud/go-http](https://github.com/sacloud/go-http) を利用する。  
これにより以下のような機能がlibsacloudと共通化される。

- リトライ/バックオフ
- 認証情報(APIキー)の取り扱い
- ロギング(http.RoundTripper)

### `service`インターフェースの実装

CLIである [Usacloud](https://github.com/sacloud/usacloud) で利用するために、`service`インターフェースを提供したい。

> `service`インターフェースとは、CRUD+L操作を行うための窓口となるインターフェースで、libsacloudでは [helper/serviceパッケージ](https://pkg.go.dev/github.com/sacloud/libsacloud/v2@v2.27.1/helper/service) で提供されている。  
> 将来的にメタデータを持たせ、AWSでいう [Cloud Control API](https://aws.amazon.com/jp/cloudcontrolapi/) 的な役割を担わせる予定だがlibsacloud v2時点ではメタデータなしでGoのコードのみ提供されている。

serviceインターフェースを用意することでUsacloudからは処理の大部分を自動生成させることができる。

## レポジトリ

 - GitHub: [https://github.com/sacloud/phy-go](https://github.com/sacloud/phy-go)

## メンバー

- [Kazumichi Yamamoto(@yamamoto-febc)](https://github.com/yamamoto-febc)

# 履歴

- 2021/11/9 @yamamoto-febc 初版作成