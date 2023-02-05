// @author cold bin
// @date 2023/2/5

package _redis

import (
	"context"
	"eliminate-male-appearance-anxiety/global"
	"time"
)

func SetEx(key string, val string, exp time.Duration) error {
	return global.RDB.Set(context.Background(), global.KeyAllPrefix+key, val, exp).Err()
}

func GetEx(key string) (string, error) {
	get := global.RDB.Get(context.Background(), global.KeyAllPrefix+key)
	return get.Val(), get.Err()
}
