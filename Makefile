# Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
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

SHELL = bash
GO = go
PROTOC = protoc

PROTO_PATH = api
GOPATH ?= $(shell $(GO) env GOPATH)

.PHONY: protos
protos:
	$(PROTOC) \
		--go_out=:../.. \
		--go-grpc_out=:../.. \
		--proto_path=$(PROTO_PATH) \
		$(PROTO_PATH)/*.proto

.PHONY: deps
deps:
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: clean
clean:
	@rm pkg/ttnpb/*.pb.go

# vim: ft=make
