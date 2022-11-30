GOOS=darwin
GOARCH=arm64

.PHONY: build

GIT_COMMIT := $(shell git rev-list -1 HEAD)
DATE := $(date +"%Y/%m/%d")

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.date=$(DATBUILT_BYE) -X main.gitCommit=$(GIT_COMMIT)" .
