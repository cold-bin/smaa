// @author cold bin
// @date 2023/2/7

package middlware

import (
	"eliminate-male-appearance-anxiety/api"
	"eliminate-male-appearance-anxiety/pkg/_jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			api.ResVerifiedErr(c)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			api.ResVerifiedErr(c)
			c.Abort()
			return
		}

		mc, err := _jwt.ParesToken(parts[1])
		if err != nil {
			api.ResVerifiedErr(c)
			c.Abort()
			return
		}

		c.Set("user_id", mc.UserId)
		c.Next()
	}
}
