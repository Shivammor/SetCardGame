#!/usr/bin/env bash

set -e

# Set env for WASM build
GOOS=js GOARCH=wasm go build -o main.wasm main.go


live-server
