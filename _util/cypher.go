// @author cold bin
// @date 2023/2/7

package _util

import (
	"crypto/sha256"
	"github.com/cold-bin/goutil/general/conv"
)

func HM256(target string) string {
	hash := sha256.New()
	hash.Write(conv.QuickS2B(target))
	return conv.QuickB2S(hash.Sum(nil))
}
