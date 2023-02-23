LAST_TAG=$(shell git describe --abbrev=0 --tags)
CURR_SHA=$(shell git rev-parse --verify HEAD)
#
LDFLAGS=-ldflags "-s -w -X main.version=$(LAST_TAG)"
# LDFLAGS=-ldflags "-s -w -X main.version=0.1.0"

release:
	git tag $(tag)
	git push origin $(tag)

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-drpc_out=. --go-drpc_opt=paths=source_relative --proto_path=. actor/actor.proto

test: build
	go test ./... --race

build:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/bpmonitor cmd/main.go

run:
	go run ./...

install:
	cp bin/bpmonitor ${HOME}/.local/bin

clean:
	rm bin/bpmonitor
	rm ${HOME}/.local/bin/bpmonitor

.PHONY: run
