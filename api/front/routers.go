// @author cold bin
// @date 2023/2/5

package front

import "github.com/gin-gonic/gin"

func SignupFrontRouters(r *gin.Engine) {
	f := r.Group("front")
	{
		f.POST("user/getCode", GetCode)
	}
}
