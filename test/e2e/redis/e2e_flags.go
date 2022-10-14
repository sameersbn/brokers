package redis

import (
	"flag"

	"github.com/triggermesh/brokers/test/e2e/lib"
)

var Flags lib.TestsFlags

func InitializeRedisFlags() {
	flag.StringVar(&Flags.RedisAddress, "redis-address", "0.0.0.0:6379", "Redis server address")
	flag.StringVar(&Flags.RedisPassword, "redis-password", "", "Redis server password")
	flag.StringVar(&Flags.RedisStreamPrefix, "redis-stream-prefix", "", "Prefix for redis server streams")

	flag.Parse()
}
