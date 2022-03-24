#
# Copyright 2021-2022 The phy-api-go authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
AUTHOR          ?="The phy-api-go authors"
COPYRIGHT_YEAR  ?="2021-2022"
COPYRIGHT_FILES ?=$$(find . -name "*.go" -print | grep -v "/vendor/")

default: gen fmt set-license goimports lint test

.PHONY: all
all: dist/phy-api-go-fake-server

dist/phy-api-go-fake-server: *.go
	go build -o dist/phy-api-go-fake-server cmd/phy-api-go-fake-server/*.go

.PHONY: test
test:
	TESTACC= go test ./... $(TESTARGS) -v -timeout=120m -parallel=8 -race;

.PHONY: testacc
testacc:
	TESTACC=1 go test ./... $(TESTARGS) -v -timeout=120m -parallel=8 ;

.PHONY: tools
tools:
	npm install -g @apidevtools/swagger-cli
	go install github.com/rinchsan/gosimports/cmd/gosimports@latest
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/sacloud/addlicense@latest
	go install github.com/client9/misspell/cmd/misspell@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/v1.43.0/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.43.0

.PHONY: clean
clean:
	find . -type f -name "*_gen.go" -delete
	rm apis/v1/spec/original-swagger.yaml
	rm apis/v1/spec/swagger.json

.PHONY: gen
gen: _gen fmt goimports set-license

.PHONY: _gen
_gen: apis/v1/spec/original-swagger.yaml apis/v1/spec/swagger.json apis/v1/zz_types_gen.go apis/v1/zz_client_gen.go apis/v1/zz_server_gen.go
	go generate ./...

apis/v1/spec/original-swagger.yaml: apis/v1/spec/original-swagger.json
	swagger-cli bundle apis/v1/spec/original-swagger.json -o apis/v1/spec/original-swagger.yaml --type yaml

apis/v1/spec/swagger.json: apis/v1/spec/swagger.yaml
	swagger-cli bundle apis/v1/spec/swagger.yaml -o apis/v1/spec/swagger.json --type json

apis/v1/zz_types_gen.go: apis/v1/spec/swagger.yaml apis/v1/spec/codegen/types.yaml
	oapi-codegen -config apis/v1/spec/codegen/types.yaml apis/v1/spec/swagger.yaml

apis/v1/zz_client_gen.go: apis/v1/spec/swagger.yaml apis/v1/spec/codegen/client.yaml
	oapi-codegen -config apis/v1/spec/codegen/client.yaml apis/v1/spec/swagger.yaml

apis/v1/zz_server_gen.go: apis/v1/spec/swagger.yaml apis/v1/spec/codegen/gin.yaml
	oapi-codegen -config apis/v1/spec/codegen/gin.yaml apis/v1/spec/swagger.yaml

.PHONY: goimports gosimports
goimports: gosimports
gosimports: fmt
	gosimports -l -w .

.PHONY: fmt
fmt:
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

.PHONY: set-license
set-license:
	@addlicense -c $(AUTHOR) -y $(COPYRIGHT_YEAR) $(COPYRIGHT_FILES)

.PHONY: lint lint-go lint-def
lint: lint-go lint-def

lint-go:
	golangci-lint run ./...

lint-def:
	docker run --rm -v $$PWD:$$PWD -w $$PWD stoplight/spectral:latest lint -F warn apis/v1/spec/swagger.yaml

.PHONY: godoc
godoc:
	echo "URL: http://localhost:6060/pkg/github.com/sacloud/phy-api-go/"
	godoc -http=localhost:6060

