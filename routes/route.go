package routes

import (
	"golang-api/controllers"
	"golang-api/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Route(r *gin.Engine) {
	v1 := r.Group("")
	{
		v1.POST("user/register", controllers.RegisterUser)
		v1.POST("user/login", controllers.LoginUser)

		v1.GET("post/list", controllers.ListPost)
		v1.GET("post/detail/:id", controllers.DetailPost)
		v1.POST("post/create", controllers.CreatePost)
		v1.PUT("post/edit/:id", controllers.UpdatePost)
		v1.DELETE("post/delete/:id", controllers.DeletePost)
	}
	secured := r.Group("").Use(middleware.AuthJWT())
	{
		secured.GET("user/me", controllers.UserInfo)
		secured.PUT("user/edit", controllers.UpdateUser)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
