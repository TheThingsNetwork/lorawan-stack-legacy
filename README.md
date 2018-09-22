# The Things Network LoRaWAN Stack Legacy API and Tools

This repository holds legacy The Things Network APIs and tools, for compatibility purposes.

## Install

Make sure your [Go environment](https://github.com/TheThingsNetwork/lorawan-stack/blob/master/DEVELOPMENT.md#getting-started-with-go-development) is set up.

```bash
git clone git@github.com:TheThingsNetwork/lorawan-stack-legacy.git "$(go env GOPATH)"/go.thethings.network/lorawan-stack-legacy
```

## Build

The content of this repository is not meant to be edited, considering it is meant to stay compatible with older software.

You can find in this repository the protocol buffer files that define the API, and the generated Go files. Make sure [Docker](https://www.docker.com/) is setup on your machine. To recreate the Go files, you can use the following commands:

```bash
make go.protos.clean # Removes the pre-existing Go files
make go.protos # Creates the Go files
```
