#!/bin/bash -x

set -e

# run tests excluding some non mocked stuff
go test -coverprofile .testCoverage.txt

status=$?
if [ $status -eq 0 ]
then
  echo "Tests Passed."
else
  rm -rf *
  echo "Error: Tests Failed."
  exit $status
fi

# install staticcheck
go install honnef.co/go/tools/cmd/staticcheck@latest

# export gobin path
export PATH="$PATH:$(go env GOPATH)/bin"

# run staticcheck in current path and exit with error if not passed
staticcheck ./...

status=$?
if [ $status -eq 0 ]
then
  echo "Staticcheck successful."
else
  echo "Error: Staticcheck Failed."
  exit $status
fi

env GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o graphql main.go

status=$?
## get status ##
## take some decision ##
if [ $status -eq 0 ]
then
  echo "Compilation succeeded."
else
  rm -rf *
  echo "Error: Compilation Failed."
  exit $status
fi