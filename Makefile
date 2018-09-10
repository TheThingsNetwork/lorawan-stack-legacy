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

go.protos:
	@which docker > /dev/null || (echo "Please install Docker to generate protos." && exit 1)
	docker run --user `id -u` --rm -v $(PWD):$(PWD) -w $(PWD) thethingsindustries/protoc:3.0.9 -I$(PWD) --gogottn_out=Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,plugins=grpc:$(GOPATH)/src discovery/discovery.proto
	docker run --user `id -u` --rm -v $(PWD):$(PWD) -w $(PWD) thethingsindustries/protoc:3.0.9 -I$(PWD) --gogottn_out=Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,plugins=grpc:$(GOPATH)/src router/router.proto

go.protos.clean:
	@for pkg in discovery router; do rm "$${pkg}/$${pkg}.pb.go"; done
