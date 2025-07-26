FROM golang:1.22 AS builder
WORKDIR /app
COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gha-linter-optimizer ./cmd/main.go

FROM debian:bookworm-slim
COPY --from=builder /app/gha-linter-optimizer /gha-linter-optimizer
ENTRYPOINT ["/gha-linter-optimizer"]
