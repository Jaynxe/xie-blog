package core

import (
	"context"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	ctx := context.Background()
	Conf := global.GVB_CONFIG.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     Conf.Addr(),
		Password: Conf.Password,
		DB:       Conf.DBName,
		PoolSize: Conf.PoolSize,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.GVB_LOGGER.Fatalf("redis connect failure:%s", err.Error())
		return
	}
	global.GVB_REDIS = rdb
	global.GVB_LOGGER.Info("redis connect succeed")
}
