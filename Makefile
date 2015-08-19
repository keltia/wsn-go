# Main Makefile for surv-export

VPATH=  surv-export:config:surv
DEST=   bin

all:    ${DEST}/surv-export

clean:
	rm -f ${DEST}/surv-export

install:
	go install -v surv-export/surv-export.go surv-export/cli.go

${DEST}/surv-export:    surv-export.go config.go client.go types.go cli.go server.go
	go build -v -o $@ surv-export/surv-export.go surv-export/cli.go

push:
	git push --all
	git push --all backup
	git push --all bitbucket
