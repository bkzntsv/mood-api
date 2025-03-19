FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080

CMD ["./main"] 