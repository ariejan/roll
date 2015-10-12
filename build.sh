#!/usr/bin/env sh

go build -ldflags "-X github.com/ariejan/roll/core.Build=`git rev-parse HEAD`" -o roll main.go
