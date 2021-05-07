FROM golang:latest

WORKDIR /app
COPY go.mod go.sum main.go /app/
RUN go build .