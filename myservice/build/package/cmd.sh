#!/bin/sh

SERVICE=myservice

rebuild()
{
    killall -q $SERVICE
    go install cmd/$SERVICE \
        && go test pkg/... \
        && $GOPATH/bin/$SERVICE &
}

rebuild

while inotifywait -e modify -rqq /go/src/; do rebuild; done

exit 1
