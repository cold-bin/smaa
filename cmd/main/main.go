// @author cold bin
// @date 2023/2/3

package main

import (
	"eliminate-male-appearance-anxiety/api"
	"eliminate-male-appearance-anxiety/conf"
	"eliminate-male-appearance-anxiety/dao/_mysql"
	"eliminate-male-appearance-anxiety/dao/_redis"
	"eliminate-male-appearance-anxiety/global"
	"eliminate-male-appearance-anxiety/pkg/oss"
	"eliminate-male-appearance-anxiety/pkg/sms"
	"github.com/cold-bin/goutil/general/mlog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	api.SignupAllRouters(r)

	if err := r.Run(":" + global.Set.App.Port); err != nil {
		panic(err)
	}
}

func init() {
	if err := conf.InitConf(); err != nil {
		panic(err)
	}
	mlog.Debug("全局配置加载成功...")
	oss.InitOss()
	mlog.Debug("七牛云oss加载成功...")
	if err := sms.InitSmsClient(); err != nil {
		panic(err)
	}
	mlog.Debug("短信客户端配置成功...")
	if err := _mysql.InitMysql(); err != nil {
		panic(err)
	}
	mlog.Debug("mysql加载成功...")
	if err := _redis.InitRedis(); err != nil {
		panic(err)
	}
	mlog.Debug("redis加载成功...")
}
