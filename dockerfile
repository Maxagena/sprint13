FROM golang:1.25.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o todo-app .


FROM ubuntu:latest

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/todo-app .

COPY --from=builder /app/web ./web

ENV TODO_PORT=7540 \
    TODO_DBFILE=/app/scheduler.db

EXPOSE $TODO_PORT
CMD ["./todo-app"]