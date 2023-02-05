// @author cold bin
// @date 2023/2/5

package service

import (
	"eliminate-male-appearance-anxiety/_util"
	"eliminate-male-appearance-anxiety/dao/_redis"
	"eliminate-male-appearance-anxiety/global"
	"eliminate-male-appearance-anxiety/model/_req"
	"eliminate-male-appearance-anxiety/pkg/sms"
	"time"
)

func GetCode(param _req.GetCode) (err error) {
	code := _util.Rand6Code()
	if err = sms.Send(param.Phone, code); err != nil {
		return
	}
	// 存储redis
	return _redis.SetEx(global.KeyCodePrefix, code, time.Minute*5)
}
