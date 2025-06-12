# syntax=docker/dockerfile:1
FROM golang:1.24.2-alpine3.21 as base

ENV DOCKER_CONTENT_TRUST=1

FROM base as build
WORKDIR /app
COPY . /app/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o servers-apis \
/app/cmd/orders/main.go /app/cmd/orders/wire_gen.go

WORKDIR /app
ENTRYPOINT [ "./servers-apis" ]