# Main Makefile for surv-export

GO=		go

VPATH=  soap
SRCS=   client.go data.go  pull.go push.go push_internal.go push_server.go \
	topic.go types.go config.go \
	operations.go templates.go

all: build

clean:
	${GO} clean -v

demo:
	${GO} build -v ./cmd/...

build: ${SRCS}
	${GO} build -v ./...

test: ${SRCS}
	${GO} test -v ./...

build: ${SRCS}
	${GO} build -v ./...

lint:
	gometalinter ./...

push:
	git push --all
	git push --tags
