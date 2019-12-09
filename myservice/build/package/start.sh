#!/bin/sh

set -o errexit

go install cmd/myservice

go test pkg/...

$GOPATH/bin/myservice
