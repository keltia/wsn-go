# Main Makefile for surv-export

SRC=	client.go config.go data.go pull.go push.go \
	soap.go topic.go xml.go

all: build test

clean:
	go clean -v

demo:
	go build -v ./cmd/...

build: ${SRC}
	go build -v ./...

test: ${SRC}
	go test -v ./...

push:
	git push --all
	git push --tags
