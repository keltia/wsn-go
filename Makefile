# Main Makefile for surv-export

VPATH=  surv-export:config:wsn

all: surv-export

clean:
	rm -f surv-export

install:
	go install -v surv-export/surv-export.go surv-export/cli.go

surv-export:    surv-export.go config.go client.go types.go cli.go server.go
	go build -v -o $@ surv-export/surv-export.go surv-export/cli.go

push:
	git push --all
	git push --all gitlab
