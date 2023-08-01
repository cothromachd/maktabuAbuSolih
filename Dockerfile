FROM golang:1.20.4-alpine as builder

WORKDIR /usr/app/maktabu_bot

RUN apk update && apk upgrade && \
    apk add --no-cache git

COPY ./go.* ./

RUN go mod download

COPY . .

RUN go build -o .bin/main main.go

FROM alpine

WORKDIR /usr/app/maktabu_bot

RUN apk update && apk upgrade && \
    apk add --no-cache git bash

COPY --from=builder /usr/app/maktabu_bot/.bin/main ./.bin/main
COPY --from=builder /usr/app/maktabu_bot/migrations ./migrations

CMD [".bin/main"]