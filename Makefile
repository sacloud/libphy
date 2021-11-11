#
# Copyright 2021 The phy-go authors
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
AUTHOR          ?="The phy-go authors"
COPYRIGHT_YEAR  ?="2021"
COPYRIGHT_FILES ?=$$(find . -name "*.go" -print | grep -v "/vendor/")

default: gen fmt set-license goimports lint test

.PHONY: test
test:
	TESTACC= go test ./... $(TESTARGS) -v -timeout=120m -parallel=8 -race;

.PHONY: testacc
testacc:
	TESTACC=1 go test ./... $(TESTARGS) -v -timeout=120m -parallel=8 ;

.PHONY: tools
tools:
	npm install -g @apidevtools/swagger-cli
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/sacloud/addlicense@latest
	go install github.com/client9/misspell/cmd/misspell@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/v1.43.0/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.43.0

.PHONY: clean
clean:
	find . -type f -name "*_gen.go" -delete
	rm openapi/spec/original-swagger.yaml
	rm openapi/spec/swagger.json

.PHONY: gen
gen: _gen fmt goimports set-license

.PHONY: _gen
_gen: openapi/spec/original-swagger.yaml openapi/spec/swagger.json openapi/zz_types_gen.go openapi/zz_client_gen.go openapi/zz_server_gen.go
	go generate ./...

openapi/spec/original-swagger.yaml: openapi/spec/original-swagger.json
	swagger-cli bundle openapi/spec/original-swagger.json -o openapi/spec/original-swagger.yaml --type yaml

openapi/spec/swagger.json: openapi/spec/swagger.yaml
	swagger-cli bundle openapi/spec/swagger.yaml -o openapi/spec/swagger.json --type json

openapi/zz_types_gen.go: openapi/spec/swagger.yaml
	oapi-codegen -generate=types -package openapi -o openapi/zz_types_gen.go openapi/spec/swagger.yaml

openapi/zz_client_gen.go: openapi/spec/swagger.yaml
	oapi-codegen -generate=client -package openapi -o openapi/zz_client_gen.go openapi/spec/swagger.yaml

openapi/zz_server_gen.go: openapi/spec/swagger.yaml
	oapi-codegen -generate=gin -package openapi -o openapi/zz_server_gen.go openapi/spec/swagger.yaml

.PHONY: goimports
goimports: fmt
	goimports -l -w .

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
	docker run --rm -v $$PWD:$$PWD -w $$PWD stoplight/spectral:latest lint -F warn openapi/spec/swagger.yaml
