FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./backend/main.go

EXPOSE 5174

CMD ["./main"]