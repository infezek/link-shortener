FROM golang:latest

RUN go get github.com/canthefason/go-watcher/cmd/watcher

WORKDIR /go/src/github.com/app/go_watcher

COPY . .

ENTRYPOINT /go/bin/watcher

