# syntax=docker/dockerfile:1
FROM golang:latest as build

ADD . /app
WORKDIR /app
# Run Build binary
RUN go build -v -o ./src/cmd/main ./src/cmd
EXPOSE 8080
CMD ["/app/src/cmd/main"]