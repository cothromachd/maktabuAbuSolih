FROM golang:alpine

WORKDIR /maktabu_bot

COPY . .

RUN go mod download
RUN go build -o main .


CMD ["./main"]