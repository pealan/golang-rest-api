FROM golang:1.23.4-bullseye AS base
WORKDIR /src
COPY go.* ./
RUN go mod download

FROM base AS build
COPY . .
RUN go build -o /server .
EXPOSE 8080
ENTRYPOINT ["/server"]

FROM base AS unit-test
COPY . .
RUN go test -v ./...
