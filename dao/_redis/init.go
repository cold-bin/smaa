// @author cold bin
// @date 2023/2/4

package _redis

import (
	"context"
	"eliminate-male-appearance-anxiety/global"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitRedis() error {
	_redis := global.Set.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", _redis.Host, _redis.Port),
		Password: _redis.Password,
		DB:       _redis.Db,
	})

	res, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	if res != "PONG" {
		return errors.New("redis link: the result of PING is not PONG")
	}

	global.RDB = rdb
	return nil
}
