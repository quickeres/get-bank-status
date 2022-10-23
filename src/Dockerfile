# Create a docker file for building a compiler file
# And save into local directory
FROM golang:1.11.1 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...
