# Build stage
FROM golang:1.23 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./app/main.go

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates netcat-openbsd
WORKDIR /root/
COPY --from=builder /app/myapp .
CMD ["sh", "-c", "until nc -z db 3306; do sleep 1; done; ./myapp"]
