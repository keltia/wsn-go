# Main Makefile for surv-export

VPATH=  surv-export:config:surv
DEST=   bin

all:    ${DEST}/surv-export

clean:
	rm -f ${DEST}/surv-export

${DEST}/surv-export:    main.go config.go client.go types.go
	go build -v -o $@ surv-export/main.go

