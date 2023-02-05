// @author cold bin
// @date 2023/2/5

package oss

import (
	"eliminate-male-appearance-anxiety/global"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func InitOss() {
	global.QBox = qbox.NewMac(global.Set.Oss.AccessKey, global.Set.Oss.SecretKey)
}

func Token() (token string) {
	ossConf := global.Set.Oss

	putPolicy := storage.PutPolicy{
		Scope:      ossConf.Bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)"}`,
	}

	putPolicy.Expires = ossConf.UploadTokenExp
	return putPolicy.UploadToken(global.QBox)
}
