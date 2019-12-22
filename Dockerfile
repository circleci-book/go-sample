FROM golang:1.13.3

WORKDIR /go/src/project

COPY . .

CMD go run main.go
