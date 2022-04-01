# syntax=docker/dockerfile:1

################################################################################
# Golang stage
################################################################################
FROM docker.io/library/golang:alpine AS golang

SHELL ["/bin/ash", "-euxo", "pipefail", "-c"]

# hadolint ignore=DL3018
RUN apk add --no-cache --upgrade \
    bc \
    binutils-gold \
    curl \
    gcc \
    git \
    git-lfs \
    gnupg \
    make \
    musl-dev

################################################################################
# Golangci stage
################################################################################
FROM golang AS golangci

RUN go install -trimpath -ldflags "-extldflags -static" \
    github.com/golangci/golangci-lint/cmd/golangci-lint@latest

################################################################################
# Gotestsum stage
################################################################################
FROM golang AS gotestsum

RUN go install -trimpath -ldflags "-extldflags -static" \
    gotest.tools/gotestsum@latest

################################################################################
# Tools stage
################################################################################
FROM golang AS tools

COPY --from=golangci /go/bin /usr/local/bin/
COPY --from=gotestsum /go/bin /usr/local/bin/

COPY ./scripts /usr/local/bin/

RUN chmod +x /usr/local/bin/*
