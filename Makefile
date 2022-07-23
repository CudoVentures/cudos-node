PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags))
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=cudos-node \
	-X github.com/cosmos/cosmos-sdk/version.AppName=cudos-noded \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

install: go.sum
		@echo "--> Installing cudos-noded"
		@go install -mod=readonly $(BUILD_FLAGS) -tags "ledger" ./cmd/cudos-noded

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)

###############################################################################
###                                Protobuf                                 ###
###############################################################################

proto-all:  proto-gen #proto-lint proto-gen TODO - Add format and lint

proto-gen:
	@echo "Generating Protobuf files"
	sh ./scripts/protocgen;


