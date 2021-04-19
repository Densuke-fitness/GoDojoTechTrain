FROM golang:1.14-alpine3.11 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/api -ldflags '-s -w'

FROM alpine:3.11
COPY --from=builder /app/bin/api /app/api

EXPOSE 8080

WORKDIR /app/
ENTRYPOINT ["/app/api"]
