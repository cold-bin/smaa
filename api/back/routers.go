// @author cold bin
// @date 2023/2/5

package back

import "github.com/gin-gonic/gin"

func SignupBackRouters(r *gin.Engine) {
	b := r.Group("back")
	{
		b.POST("")
	}
}
