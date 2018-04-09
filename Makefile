# Main Makefile for surv-export

VPATH=  config:soap:wsn
SRCS=   client.go data.go  pull.go push.go push_server.go topic.go types.go \
	config.go \
	operations.go templates.go

all: build

clean:
	go clean -v

build: ${SRCS}
	go build -v ./...
	go test -v ./...

lint:
	gometalinter ./...

push:
	git push --all
	git push --all gitlab
