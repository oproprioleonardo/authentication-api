FROM golang:1.20-alpine AS builder

WORKDIR /usr/src/app

COPY . .
RUN go get -d -v ./...
RUN go build -v
FROM alpine:latest

# We'll likely need to add SSL root certificates
RUN apk --no-cache add ca-certificates

WORKDIR /usr/local/bin
COPY --from=builder /usr/src/app/.env .env
COPY --from=builder /usr/src/app/privateapi .
RUN chmod 777 .env
RUN chmod +x privateapi
CMD ["./privateapi"]
