#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/github.com/maibornwolff/vbump
COPY . .
RUN apk add --no-cache -v git gcc libc-dev
RUN go get -v ./... && go get "github.com/onsi/gomega"
RUN go test -v
RUN go build
#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/maibornwolff/vbump/vbump /vbump
RUN mkdir /data
ENTRYPOINT ./vbump -d data
LABEL Name=vbump Version=1.0.0
EXPOSE 8080
