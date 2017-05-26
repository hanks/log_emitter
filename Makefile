BINARY=log_emitter

VERSION=1.0.0
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.DEFAULT_GOAL := build

test:
	cd log_emitter && go test

benchmark:
	cd log_emitter && go test -test.bench=".*"

build:
	go fmt
	cd log_emitter && go test
	mkdir -p bin/
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY}.linux64
	go build ${LDFLAGS} -o bin/${BINARY}.macos

clean:
	cd bin/
	if [ -f ${BINARY}.macos ]; then rm ${BINARY}.macos ; fi
	if [ -f ${BINARY}.linux64 ]; then rm ${BINARY}.linux64 ; fi
        
.PHONY: clean, build

