FROM golang:1.9.0-alpine

WORKDIR /go/src/github.com/lucj/who

COPY http.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o http .

CMD ["./http"]
