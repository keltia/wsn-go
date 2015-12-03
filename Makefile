# Main Makefile for surv-export

VPATH=  config:wsn

all: test

clean:
	rm -f surv-export

test: client.go server.go types.go config.go
	go test -v ./...

push:
	git push --all
	git push --all backup
