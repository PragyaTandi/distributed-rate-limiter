# Distributed Rate Limiter

A distributed rate limiter built in Go using the Token Bucket algorithm.

## Features

- Token Bucket rate limiting
- HTTP middleware
- In-memory implementation
- Thread-safe using sync.Mutex
- HTTP 429 responses
- Unit tests
- Benchmark tests

## Project Structure

```
cmd/
internal/
tests/
```

## Run

```bash
go run cmd/server/main.go
```

Server:

```
http://localhost:8080
```

## Run Tests

```bash
go test ./...
```

## Run Benchmark

```bash
go test ./tests -bench=BenchmarkTokenBucket -run=^$ -count=1
```