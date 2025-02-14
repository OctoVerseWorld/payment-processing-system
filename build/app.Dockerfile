FROM golang:1.21 AS builder

WORKDIR /app

COPY ../deployments .

RUN go mod tidy

RUN go build -o payments_service .

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/payments_service .

EXPOSE 8080

CMD ["/app/payments_service"]