FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.work go.work.sum ./
COPY apps/api-gateway/go.mod apps/api-gateway/go.sum ./apps/api-gateway/
COPY apps/execution-service/go.mod apps/execution-service/go.sum ./apps/execution-service/
COPY pkg/go.mod pkg/go.sum ./pkg/
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bytelearn-execution ./apps/execution-service/cmd/main.go

FROM alpine:3.19
WORKDIR /app
RUN apk --no-cache add ca-certificates tzdata docker-cli
COPY --from=builder /bytelearn-execution .
EXPOSE 8081
CMD ["./bytelearn-execution"]