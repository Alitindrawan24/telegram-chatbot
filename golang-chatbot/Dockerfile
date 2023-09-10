FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY . /app

RUN go build -v -o golang-chatbot main.go

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/golang-chatbot /app/golang-chatbot
COPY --from=builder /app/.env /app/.env

EXPOSE 8080

CMD ["./golang-chatbot"]
