// @author cold bin
// @date 2023/2/5

package api

import (
	"eliminate-male-appearance-anxiety/api/back"
	"eliminate-male-appearance-anxiety/api/front"
	"github.com/gin-gonic/gin"
)

func SignupAllRouters(r *gin.Engine) {
	front.SignupFrontRouters(r)
	back.SignupBackRouters(r)
}
