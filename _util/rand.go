// @author cold bin
// @date 2023/2/5

package _util

import (
	"fmt"
	"math/rand"
	"time"
)

// Rand6Code 随机生产6位随机数
func Rand6Code() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
