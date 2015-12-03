# Main Makefile for surv-export

VPATH=  config:wsn

all: test

clean:
	go clean -v

test: client.go server.go types.go config.go
	go test -v ./...

push:
	git push --all
	git push --all backup
