local key = KEYS[1]

local capacity = tonumber(ARGV[1])
local refillRate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])


local bucket = redis.call("HMGET", key, "tokens", "timestamp")


local tokens = tonumber(bucket[1])
local timestamp = tonumber(bucket[2])


if tokens == nil then
    tokens = capacity
    timestamp = now
end


local elapsed = now - timestamp

tokens = math.min(
    capacity,
    tokens + elapsed * refillRate
)


local allowed = 0


if tokens >= 1 then
    tokens = tokens - 1
    allowed = 1
end


redis.call(
    "HMSET",
    key,
    "tokens",
    tokens,
    "timestamp",
    now
)


redis.call(
    "EXPIRE",
    key,
    3600
)


return allowed