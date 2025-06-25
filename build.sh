#!/usr/bin/env bash

cd /home/vscode/github.com/si-arakaki/account-manager \
 && GOOS=darwin GOARCH=arm64 go build -o ./dist/account ./cmd/account
