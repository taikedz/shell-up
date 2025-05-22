#!/usr/bin/env bash

set -euo pipefail

THIS="$(readlink -f "$0")"
HERE="$(dirname "$THIS")"

cd "$HERE"

mkdir -p bin
go build -o bin/shup shup.go
