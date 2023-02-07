// @author cold bin
// @date 2023/2/5

package api

import (
	"eliminate-male-appearance-anxiety/api/back"
	"eliminate-male-appearance-anxiety/api/front"
	"eliminate-male-appearance-anxiety/api/middlware"
	"github.com/gin-gonic/gin"
)

func SignupAllRouters(r *gin.Engine) {
	r.Use(middlware.Cors())
	front.SignupFrontRouters(r)
	back.SignupBackRouters(r)
}
