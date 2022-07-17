package db

import (
	"shape-api/config"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
	var redisHost = config.GetConfig("redis.host")
	var redisPassword = config.GetConfig("redis.password")
	var redisDatabase, _ = strconv.Atoi(config.GetConfig("redis.database"))

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       redisDatabase,
	})
}
