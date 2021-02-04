FROM golang:latest

RUN mkdir -p /go/src/github.com/rahmancam/go-quote-api

WORKDIR /go/src/github.com/rahmancam/go-quote-api

COPY . /go/src/github.com/rahmancam/go-quote-api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/rahmancam/go-quote-api

EXPOSE 8080

ENTRYPOINT ["/go/bin/go-quote-api"]