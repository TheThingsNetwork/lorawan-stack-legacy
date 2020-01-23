# The Things Network Stack V2 Legacy API

This repository contains legacy The Things Network V2 API for compatibility purposes.

## Install

Make sure your [Go environment](https://github.com/TheThingsNetwork/lorawan-stack/blob/master/DEVELOPMENT.md#getting-started-with-go-development) is set up.

```bash
$ git clone git@github.com:TheThingsNetwork/lorawan-stack-legacy.git $(go env GOPATH)/go.thethings.network/lorawan-stack-legacy
$ cd $(go env GOPATH)/go.thethings.network/lorawan-stack-legacy
```

Also, make sure you have [`protoc`](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation) installed.

## Regenerate

The contents of this repository is not meant to be edited, considering it is meant to stay compatible with older software.

You can find in this repository the protocol buffer files that define the API, and the generated Go files.

To regenerate the Go files, use the following commands:

```bash
$ make deps   # Get dependencies
$ make clean  # Removes generated protos
$ make protos # Generates the protos
```
