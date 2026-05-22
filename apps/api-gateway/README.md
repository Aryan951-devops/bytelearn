# ByteLearn API Gateway

## Overview

The API Gateway is the central backend service of ByteLearn responsible for handling authentication, business logic, database operations, and communication between platform services.

The backend is designed using modular and production-inspired architecture principles to support scalability, maintainability, and future microservice integration.

---

# Running the Server

## Install Dependencies

```bash
go install golang.org/x/tools/cmd/goimports@latest
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
go install github.com/pressly/goose/v3/cmd/goose@latest
go mod tidy
```

## Creating Tables

```bash
goose -dir migrations postgres "<supabase_url>" up
```

## Deleting Tables
```bash
goose -dir migrations postgres "<supabase_url>" down
```

## Start Server

```bash
go run cmd/server/main.go
```

---