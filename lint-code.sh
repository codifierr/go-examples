#!/bin/bash -e

golangci-lint run ./...

pre-commit run --all-files
