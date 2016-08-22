FROM golang:1.7
WORKDIR /go/src/github.com/therealplato/muid
ADD *.go ./
CMD go run example/main.go

