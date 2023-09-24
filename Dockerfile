FROM scratch as scratch
WORKDIR /app

COPY ./common ./common
COPY ./types ./types
COPY ./input.json ./input.json

FROM golang:1.21.1-alpine3.18 as base

WORKDIR /app

COPY go.mod go.sum ./
COPY main.go ./

RUN go mod download

COPY --from=scratch /app/input.json ./
COPY --from=scratch /app/common ./common
COPY --from=scratch /app/types ./types


ENTRYPOINT [ "go", "run", "main.go" ]

