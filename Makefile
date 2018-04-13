# Main Makefile for surv-export

VPATH=  soap
SRCS=   client.go data.go  pull.go push.go push_internal.go push_server.go \
	topic.go types.go config.go \
	operations.go templates.go

all: build

clean:
	go clean -v

demo:
	go build -v ./cmd/...

build: ${SRCS}
	go build -v ./...

test: ${SRCS}
	go test -v ./...

lint:
	gometalinter ./...

push:
	git push --all
	git push --tags
