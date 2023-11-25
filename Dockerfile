FROM golang:1.20.6-alpine3.17 AS build

WORKDIR /go/src/app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod tidy
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./cmd/main.go


FROM alpine:3.17

WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin

ENTRYPOINT /go/bin/web-app
