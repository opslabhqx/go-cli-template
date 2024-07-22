# syntax=docker/dockerfile:1

# renovate: datasource=docker depName=alpine
ARG ALPINE_VERSION=3.20.1
# renovate: datasource=docker depName=golang
ARG GO_VERSION=1.22.4

FROM alpine:${ALPINE_VERSION} AS base

# Base image for the final stage
RUN apk add --no-cache bash git curl go aws-cli

FROM golang:${GO_VERSION}-alpine AS build

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go mod tidy \
    && CGO_ENABLED="0" go build \
    -ldflags="-X 'go-cli/app/internal/config/base.APP_NAME=go-cli' -X 'go-cli/app/internal/config/base.APP_VERSION=$CI_COMMIT_REF_NAME'" \
    -o go-cli **/*.go

FROM base

WORKDIR /app

COPY --from=build /app/go-cli .
