FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o golden-path-cli main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/golden-path-cli .
COPY --from=builder /app/templates ./templates
ENTRYPOINT ["./golden-path-cli"]
