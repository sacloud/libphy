#
# Copyright 2021 The libphy authors
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
AUTHOR          ?="The libphy authors"
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
	rm definitions/swagger.yaml

.PHONY: gen
gen: _gen fmt goimports set-license

.PHONY: _gen
_gen: definitions/original-swagger.yaml
	oapi-codegen -generate=types -package phy -o zz_types_gen.go definitions/swagger.yaml
	oapi-codegen -generate=client -package phy -o zz_client_gen.go definitions/swagger.yaml
	go generate ./...

definitions/original-swagger.yaml: definitions/original-swagger.json
	swagger-cli bundle definitions/original-swagger.json -o definitions/original-swagger.yaml --type yaml

.PHONY: goimports
goimports: fmt
	goimports -l -w .

.PHONY: fmt
fmt:
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: set-license
set-license:
	@addlicense -c $(AUTHOR) -y $(COPYRIGHT_YEAR) $(COPYRIGHT_FILES)

.PHONY: lint_definition
lint_definition:
	docker run -it --rm -v $$PWD:$$PWD -w $$PWD stoplight/spectral:latest lint definitions/*.{json,yaml,yml}
