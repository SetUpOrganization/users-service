FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o ./bin/run_app ./cmd/main/main.go

FROM alpine:3.20 AS runner
WORKDIR /app

COPY --from=builder /app/internal/infrastructure/db/migrations/ ./migrations
COPY --from=builder /app/bin ./

# Install goose for db migrations
RUN wget -O /usr/local/bin/goose https://github.com/pressly/goose/releases/download/v3.24.1/goose_linux_x86_64
RUN chmod +x /usr/local/bin/goose

# Migrate db and run server
CMD ["sh", "-c", "goose -dir ./migrations postgres $DATABASE_URL up && ./run_app"]