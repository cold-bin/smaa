// @author cold bin
// @date 2023/2/5

package front

import (
	"eliminate-male-appearance-anxiety/api/front/handle"
	"eliminate-male-appearance-anxiety/api/middlware"
	"github.com/gin-gonic/gin"
)

func SignupFrontRouters(r *gin.Engine) {
	f := r.Group("front")
	{
		fu := f.Group("user")
		fu.POST("get-code", handle.GetCode)
		fu.POST("login-by-code", handle.LoginByCode)
		fu.POST("login-by-password", handle.LoginByPassword)
		fu.POST("register", handle.Register)
		fu.POST("reset-password", handle.RegretPassword)

		f.Use(middlware.Jwt())

		f.GET("oss/token", handle.GetOssToken) // todo

		f.POST("daily-question", handle.GetDailyQuestion)
		f.GET("work/hot", handle.GetHot)
		f.GET("work", handle.GetWork)
		f.POST("work", handle.CreateWork)
		f.GET("search", handle.Search)

		f.POST("focus", handle.Focus)
		f.GET("focus", handle.GetFocuses)
		f.GET("focus/work", handle.GetFocusWork)

		f.POST("comment", handle.CreateComment)
		f.GET("comment", handle.GetComment)

		f.POST("like", handle.Like)
		f.POST("collection", handle.Collect)

		fu.GET("self", handle.GetSelf)
		fu.GET("self/work", handle.GetSelfWork)
		fu.GET("self/collection", handle.GetSelfCollection)
		fu.GET("self/like", handle.GetSelfLike)

		fu.GET("other", handle.GetOther)
		fu.GET("other/work", handle.GetOtherWork)
		fu.GET("other/collection", handle.GetOtherCollection)
		fu.GET("other/like", handle.GetOtherLike)

		fm := f.Group("message")
		{
			fm.GET("like", handle.GetLikeMsg)
			fm.GET("focus", handle.GetFocusMsg)
			fm.GET("collection", handle.GetCollectionMsg)
			fm.GET("comment", handle.GetCommentMsg)
		}

	}
}
