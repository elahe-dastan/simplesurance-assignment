# syntax=docker/dockerfile:experimental
FROM golang:alpine as builder

# git is an optional requirement for go since 1.18.
RUN apk add git

WORKDIR /app

# project doesn't have external dependencies so these steps are
# redundant.
# COPY go.mod go.sum ./
# RUN go mod download

COPY . .
RUN go build -o /simplesurance-assignment

FROM alpine:latest as release

WORKDIR /app/

COPY --from=builder /simplesurance-assignment .

EXPOSE 1378
ENTRYPOINT ["./simplesurance-assignment"]