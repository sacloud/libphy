#
# Copyright 2021-2025 The phy-api-go authors
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
#====================
AUTHOR         ?= The sacloud/phy-api-go authors
COPYRIGHT_YEAR ?= 2021-2025

BIN            ?= dist/phy-api-go-fake-server
GO_ENTRY_FILE  ?= cmd/phy-api-go-fake-server/*.go
BUILD_LDFLAGS  ?=

include includes/go/common.mk
include includes/go/single.mk
#====================

default: gen $(DEFAULT_GOALS)

.PHONY: tools
tools: dev-tools
	npm install -g @apidevtools/swagger-cli
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.2


.PHONY: clean-all
clean-all:
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
	oapi-codegen --old-config-style -config apis/v1/spec/codegen/types.yaml apis/v1/spec/swagger.yaml

apis/v1/zz_client_gen.go: apis/v1/spec/swagger.yaml apis/v1/spec/codegen/client.yaml
	oapi-codegen --old-config-style -config apis/v1/spec/codegen/client.yaml apis/v1/spec/swagger.yaml

apis/v1/zz_server_gen.go: apis/v1/spec/swagger.yaml apis/v1/spec/codegen/gin.yaml
	oapi-codegen --old-config-style -config apis/v1/spec/codegen/gin.yaml apis/v1/spec/swagger.yaml

lint-def:
	docker run --rm -v $$PWD:$$PWD -w $$PWD stoplight/spectral:latest lint -F warn apis/v1/spec/swagger.yaml
