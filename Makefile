BUILD_DIR = ./bin
WEB_DIR = ./web/build/
TIRELEASE_BINARY = ${BUILD_DIR}/tirelease

build.web:
	cd web && \
	yarn build

build.server:
	go build -o ${TIRELEASE_BINARY} cmd/tirelease/*.go

build: build.web build.server

run: build
	./${TIRELEASE_BINARY}

clean:
	rm -rf ${WEB_DIR}
	rm -rf ${BUILD_DIR}

help:
	@echo "make build : build binary for tirelease"
	@echo "make clean : clean all binary bin dictionary"

.PHONY: build.web build.server build run clean help

