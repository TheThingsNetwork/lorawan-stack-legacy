# Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

GOPATH ?= $(shell go env GOPATH)
PWD ?= $(shell pwd)

EMPTY :=
SPACE := $(EMPTY) $(EMPTY)
COMMA := ,

dev-deps:
	go get -u ./... github.com/alecthomas/gometalinter
	gometalinter -i

DOCKER ?= docker

PROTOC_OUT ?= /out

PROTOC_DOCKER_IMAGE ?= thethingsindustries/protoc:3.0.23
PROTOC_DOCKER_ARGS = run --user `id -u` --rm \
                     --mount type=bind,src=$(PWD)/api,dst=$(PWD)/api \
                     --mount type=bind,src=$(PWD)/pkg/ttnpb,dst=$(PROTOC_OUT)/go.thethings.network/lorawan-stack-legacy/pkg/ttnpb \
                     -w $(PWD)
PROTOC ?= $(DOCKER) $(PROTOC_DOCKER_ARGS) $(PROTOC_DOCKER_IMAGE) -I$(shell dirname $(PWD))

protoc:
	$(DOCKER) pull $(PROTOC_DOCKER_IMAGE)

GO_PROTO_TYPES := any duration empty field_mask struct timestamp wrappers
GO_PROTO_TYPE_CONVERSIONS = $(subst $(SPACE),$(COMMA),$(foreach type,$(GO_PROTO_TYPES),Mgoogle/protobuf/$(type).proto=github.com/gogo/protobuf/types))
GO_PROTOC_FLAGS ?= \
	--gogottn_out=plugins=grpc,$(GO_PROTO_TYPE_CONVERSIONS):$(PROTOC_OUT) \

go.protos: $(wildcard api/*.proto)
	$(PROTOC) $(GO_PROTOC_FLAGS) $(PWD)/api/*.proto
	perl -i -pe 's:golang.org/x/net/context:context:' `find ./pkg -name '*pb.go'`
	goimports -w $(PWD)/pkg/ttnpb
	unconvert -apply ./pkg/ttnpb/...
	gofmt -w -s $(PWD)/pkg/ttnpb

go.protos.clean:
	find ./pkg/ttnpb -name '*.pb.go' -delete
