BUILD_DIR = ./bin
WEB_DIR = ./web/build/
TIRELEASE_BINARY = ${BUILD_DIR}/tirelease

build.web:
	cd web && \
	yarn install && \
	yarn build

build.server:
	go build -o ${TIRELEASE_BINARY} cmd/tirelease/*.go

all: build.web build.server

run: all
	./${TIRELEASE_BINARY}

clean:
	rm -rf ${WEB_DIR}
	rm -rf ${BUILD_DIR}

help:
	@echo "make all : build all binary for tirelease"
	@echo "make clean : clean all binary bin dictionary"

.PHONY: build.web build.server all run clean help

