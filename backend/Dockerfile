FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/backend
RUN go build -o ../main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main ./
EXPOSE 5174
CMD ["./main"]
