FROM golang:1.19

ADD . /go/src/go-webserver
WORKDIR /go/src/go-webserver

RUN go install go-webserver

ENTRYPOINT /go/bin/go-webserver

EXPOSE 8080