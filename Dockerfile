FROM golang:alpine3.20 AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -o ./word-of-wisdom
 
FROM alpine:3.20.0
WORKDIR /
RUN apk add fortune
COPY --from=builder /app/word-of-wisdom .
EXPOSE 8080
CMD ["/word-of-wisdom", "run"]