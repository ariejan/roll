BINARY := roll

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

VERSION=`cat VERSION`
BUILD_DATE=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/ariejan/roll/core.Version=${VERSION} -X github.com/ariejan/roll/core.BuildDate=${BUILD_DATE}"

.DEFAULT_GOAL := $(BINARY)

.PHONY: help
help:
	echo ${VERSION}
	echo ${BUILD_DATE}

$(BINARY): $(SOURCES)
	go build ${LDFLAGS} -o ${BINARY} main.go

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
