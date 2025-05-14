FROM golang:1.23.3-alpine3.20 AS builder

ARG OS=linux
ARG ARCH=x86_64

# Install build dependencies
RUN apk add --no-cache --update \
    ca-certificates \
    git \
    make \
    build-base \
    linux-headers

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

FROM alpine:3.20

RUN apk add --no-cache ca-certificates bash jq

COPY --from=builder /app/bin/indexer /usr/local/bin/indexer

ENTRYPOINT ["indexer"]
