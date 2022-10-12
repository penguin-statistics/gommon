FROM golang:1.19.2-alpine AS base
WORKDIR /app

# tester
FROM base AS tester
ENV GOOS linux
ENV GOARCH amd64

# build-args
ARG VERSION

RUN apk --no-cache add bash git openssh

# modules: utilize build cache
COPY go.mod ./
COPY go.sum ./

# RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .

# run all go tests and generate coverage report
RUN go test -v -coverprofile=coverage.out ./...
