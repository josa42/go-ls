#!/usr/bin/env bash

declare -a files=(service.go service_test.go structures.go jsonrpc2.go)

url='https://raw.githubusercontent.com/sourcegraph/go-lsp/master'

for f in "${files[@]}"; do
  curl -sS "$url/$f" > "lsp/$f"
done

