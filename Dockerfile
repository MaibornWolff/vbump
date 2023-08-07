FROM golang:1.18-alpine as builder
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

# create non root user
RUN addgroup -S nonroot && \
    adduser -S nonroot -G nonroot

COPY --chown=nonroot:nonroot --from=builder /build/vbump /vbump

RUN mkdir /data
RUN chown -R nonroot:nonroot /data
ENTRYPOINT ./vbump -d data
LABEL Name=vbump Version=1.4.0
EXPOSE 8080

USER nonroot
