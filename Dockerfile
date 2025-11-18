FROM golang:1.22 AS builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gha-linter-optimizer ./cmd/main.go
FROM debian:bookworm-slim
RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/gha-linter-optimizer /gha-linter-optimizer
ENTRYPOINT ["/gha-linter-optimizer"]
