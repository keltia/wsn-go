# Main Makefile for surv-export

VPATH=  config:wsn

all: build test

clean:
	go clean -v

build: client.go server.go types.go config.go
	go build -v ./...

test: client.go server.go types.go config.go
	go test -v ./...

push:
	git push --all origin
	git push --all backup
