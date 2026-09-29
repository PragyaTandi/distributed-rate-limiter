local key = KEYS[1]

local capacity = tonumber(ARGV[1])
local refill_rate = tonumber(ARGV[2])
local requested_tokens = tonumber(ARGV[3])

local now = redis.call("TIME")
local current_time = tonumber(now[1])

local tokens = tonumber(redis.call("HGET", key, "tokens"))
local last_refill = tonumber(redis.call("HGET", key, "last_refill"))

if tokens == nil then
    tokens = capacity
    last_refill = current_time
end

local elapsed = current_time - last_refill
tokens = math.min(capacity, tokens + (elapsed * refill_rate))

local allowed = 0
local retry_after = 0

if tokens >= requested_tokens then
    tokens = tokens - requested_tokens
    allowed = 1
else
    if refill_rate > 0 then
        retry_after = math.ceil((requested_tokens - tokens) / refill_rate)
    else
        retry_after = -1
    end
end

redis.call("HSET", key,
    "tokens", tokens,
    "last_refill", current_time
)

return {allowed, tokens, retry_after}