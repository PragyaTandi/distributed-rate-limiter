# Distributed Rate Limiter

A high-performance distributed rate limiter built using Go.

## Features

- Token Bucket algorithm
- In-memory rate limiting
- Per-user/API-key based limiting
- Thread-safe implementation
- HTTP middleware integration
- Redis-ready architecture
- Unit tests
- Concurrent request testing
- Benchmark testing


## Architecture


## Token Bucket

The token bucket algorithm works by:

- Bucket stores tokens
- Requests consume tokens
- Tokens refill at a fixed rate
- Requests are rejected when tokens are unavailable


## Running the Server

```bash
go run cmd/server/main.go
Client
|
|
HTTP Server
|
|
Rate Limit Middleware
|
|
RateLimiter Interface
|
+----------------+
| |
Token Bucket Redis Limiter
(In Memory) (Distributed)



## Token Bucket

The token bucket algorithm works by:

- Bucket stores tokens
- Requests consume tokens
- Tokens refill at a fixed rate
- Requests are rejected when tokens are unavailable


## Running the Server

```bash
go run cmd/server/main.go

Server:

http://localhost:8080
Testing

Run all tests:

go test ./...

Run race detector:

go test ./tests -race

Run benchmark:

go test ./tests -bench=BenchmarkTokenBucket -run=^$
Example Request
curl -H "X-API-Key: user1" http://localhost:8080/
Future Improvements
Redis atomic operations
Multiple server instances
Docker Compose deployment
Load testing
Monitoring with Prometheus and Grafana

Save.

---

Then run final verification:

### PowerShell

```powershell
go test ./...