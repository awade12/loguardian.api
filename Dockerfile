# syntax=docker/dockerfile:1

FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-translation-api

FROM alpine:latest

WORKDIR /root/

COPY --from=build /go-translation-api .
COPY .env .

EXPOSE 65000

CMD ["./go-translation-api"]
