# syntax=docker/dockerfile:1

FROM golang:1.18-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o /canban
EXPOSE 8081
CMD [ "/canban" ]