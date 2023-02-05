// @author cold bin
// @date 2023/2/5

package front

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

	if err := service.GetCode(param); err != nil {
		mlog.Error(err)
	}

	api.ResOk(c)
}
