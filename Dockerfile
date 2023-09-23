FROM golang:1.21.1-alpine3.14 AS base

WORKDIR /app

COPY go.mod go.sum ./
COPY ./common ./common
COPY ./types ./types
COPY main.go ./

RUN go mod download


