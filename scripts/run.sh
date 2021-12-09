#!/usr/bin/env bash

echo "==> run project now......"

## build web
cd ./web
rm -rf build/
yarn build

## run golang server
cd ..
go run cmd/tirelease/main.go