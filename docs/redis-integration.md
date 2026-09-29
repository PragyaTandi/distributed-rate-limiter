# Redis Integration

## Redis Setup

Redis runs using Docker Compose.

Start Redis with:

```bash
docker compose up -d
```

Check that Redis is running:

```bash
docker ps
```

Test the Redis connection:

```bash
docker exec rate-limiter-redis redis-cli ping
```

Expected output:

```text
PONG
```

## Redis Key Structure

Each client gets a separate Redis key:

```text
rate_limit:{client_id}
```

The key is stored as a Redis HASH with the following fields:

```text
tokens
last_refill
```

Example:

```text
rate_limit:user123

tokens       9
last_refill  1790708607
```

## Lua Rate Limiting Script

The atomic rate-limiting logic is located at:

```text
scripts/rate_limit.lua
```

The script accepts three arguments:

```text
capacity
refill_rate
requested_tokens
```

The script:

1. Reads the client's current token state.
2. Calculates tokens regenerated since the last request.
3. Caps the token count at the bucket capacity.
4. Checks whether the request can be allowed.
5. Deducts the requested tokens when allowed.
6. Calculates `retry_after` when the request is denied.
7. Stores the updated token state.
8. Sets an expiration time on the Redis key.

The operations execute atomically because they are performed inside a Redis Lua script.

## Script Result

The script returns three values:

```text
allowed
remaining_tokens
retry_after
```

Example:

```text
1
9
0
```

This means:

* Request allowed
* 9 tokens remaining
* No retry required

A denied request may return:

```text
0
0
9978
```

This means:

* Request denied
* 0 tokens remaining
* Approximately 9978 seconds until enough tokens are available

## Manual Testing

Copy the Lua script into the Redis container:

```bash
docker cp ./scripts/rate_limit.lua rate-limiter-redis:/tmp/rate_limit.lua
```

Run the rate limiter:

```bash
docker exec rate-limiter-redis sh -c "redis-cli --eval /tmp/rate_limit.lua rate_limit:test , 10 1 1"
```

Expected output:

```text
1
9
0
```

## Redis State

Rate-limit state is stored in Redis rather than only in application memory. This allows multiple application instances to share the same rate-limit state.

Each client key automatically expires after its configured TTL when the key becomes inactive.
