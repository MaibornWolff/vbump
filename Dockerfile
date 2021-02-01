FROM golang:alpine as builder
RUN apk update && apk add build-base
RUN mkdir /build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build .
RUN go test -v ./...

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /build/vbump /vbump
RUN mkdir /data
ENTRYPOINT ./vbump -d data
LABEL Name=vbump Version=1.3.0
EXPOSE 8080
