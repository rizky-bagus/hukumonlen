FROM golang:1.20.2 AS builder
WORKDIR /app

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o deploy/api cmd/api/main.go

FROM alpine:3.13
WORKDIR /app

COPY --from=builder /app/deploy/api .

EXPOSE 8080

CMD ["./api"]