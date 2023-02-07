// @author cold bin
// @date 2023/2/5

package handle

import (
	"eliminate-male-appearance-anxiety/api"
	"eliminate-male-appearance-anxiety/model/_req"
	"eliminate-male-appearance-anxiety/service"
	"github.com/cold-bin/goutil/general/mlog"
	"github.com/cold-bin/goutil/web"
	"github.com/gin-gonic/gin"
)

func GetCode(c *gin.Context) {
	var param _req.GetCode
	if err := c.ShouldBind(&param); err != nil {
		mlog.Debug(err)
		web.ResParamErr(c)
		return
	}

	if err := service.GetCode(&param); err != nil {
		mlog.Error(err)
		api.ResInternalErr(c)
	}

	api.ResOk(c)
}

func LoginByCode(c *gin.Context) {
	var param _req.LoginByCode
	err := c.ShouldBind(&param)
	if err != nil {
		mlog.Debug(err)
		api.ResParamErr(c)
		return
	}

	res, err := service.LoginByCode(&param)
	if err != nil {

		if err == service.ErrUserPC {
			mlog.Debug(err)
			api.ResParamErr(c)
			return
		}
		mlog.Error(err)
		api.ResInternalErr(c)
		return
	}

	api.ResOkWithData(c, res)
}

func LoginByPassword(c *gin.Context) {
	var param _req.LoginByPassword
	err := c.ShouldBind(&param)
	if err != nil {
		mlog.Debug(err)
		api.ResParamErr(c)
		return
	}

	res, err := service.LoginByPassword(&param)
	if err != nil {
		if err == service.ErrPP {
			mlog.Debug(err)
			api.ResParamErr(c)
			return
		}
		mlog.Error(err)
		api.ResInternalErr(c)
		return
	}

	api.ResOkWithData(c, res)
}

func Register(c *gin.Context) {
	var param _req.Register
	err := c.ShouldBind(&param)
	if err != nil {
		mlog.Debug(err)
		api.ResParamErr(c)
		return
	}

	res, err := service.Register(&param)
	if err != nil {
		if err == service.ErrSmsCode || err == service.ErrPhone {
			mlog.Debug(err)
			api.ResParamErr(c)
			return
		}
		mlog.Error(err)
		api.ResInternalErr(c)
	}

	api.ResOkWithData(c, res)
}

func RegretPassword(c *gin.Context) {
	var param _req.RegretPassword
	err := c.ShouldBind(&param)
	if err != nil {
		mlog.Debug(err)
		api.ResParamErr(c)
		return
	}

	err = service.RegretPassword(&param)
	if err != nil {
		mlog.Error(err)
		api.ResInternalErr(c)
		return
	}

	api.ResOk(c)
}

func GetSelf(c *gin.Context) {

}

func GetOther(c *gin.Context) {

}
