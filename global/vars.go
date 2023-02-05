// @author cold bin
// @date 2023/2/4

package global

import (
	_sms "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/go-redis/redis/v8"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"gorm.io/gorm"
)

// 数据库句柄
var (
	GDB       *gorm.DB
	RDB       *redis.Client
	SmsClient *_sms.Client
	QBox      *qbox.Mac
)

var Set Settings
