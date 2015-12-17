# Main Makefile for surv-export

VPATH=  config:wsn
SRCS=   push_client.go push_server.go types.go config.go

all: build test

clean:
	go clean -v

build: ${SRCS}
	go build -v ./...

test: ${SRCS}
	go test -v ./...

push:
	git push --all
	git push --all gitlab
