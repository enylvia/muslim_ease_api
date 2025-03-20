FROM golang:1.21.3 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod build -o muslim_ease_api .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/muslim_ease_api .

CMD ["./muslim_ease_api"]