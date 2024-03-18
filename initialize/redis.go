package initialize

import (
	"strconv"

	"github.com/go-redis/redis"

	"Easy-Gin/config"
)

// InitRedis 初始化 Redis
func InitRedis() *redis.Client {
	r := config.GinConfig.Redis
	RDB := redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + strconv.Itoa(r.Port),
		Password: r.Password, // no password set
		DB:       r.DB,       // use default DB
		PoolSize: r.PoolSize, // 连接池大小
	})

	_, err := RDB.Ping().Result()
	if err != nil {
		panic(err)
	}
	return RDB
}
