# To make sure our syntax is up-to-date
# syntax=docker/dockerfile:1

FROM golang:1.17.8-alpine3.15

WORKDIR /usr/src/go-backend-demo
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/go-backend-demo .

EXPOSE 8080

CMD ["go-backend-demo"]
