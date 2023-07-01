FROM golang:1.19

ADD . /go/src/github.com/go-webserver
WORKDIR /go/src/github.com/go-webserver

RUN go install go-webserver

ENTRYPOINT /go/bin/go-webserver

EXPOSE 8080