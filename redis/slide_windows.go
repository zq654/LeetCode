package redis

const (
	slideWindows = `
local userid = KEYS[1]
local now = tonumber(ARGV[1])
local size = tonumber(ARGV[2])
local limit = tonumber(ARGV[3])

local user_set = redis.call("ZGET",userid)
`
)
