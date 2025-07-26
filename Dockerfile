FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o gha-linter ./cmd/main.go

FROM debian:bullseye-slim
COPY --from=builder /app/gha-linter-optimizer /gha-linter-optimizer
ENTRYPOINT ["/gha-linter-optimizer"]
