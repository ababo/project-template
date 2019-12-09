#!/bin/sh

set -o errexit

go install cmd/myapp

go test pkg/...

$GOPATH/bin/myapp
