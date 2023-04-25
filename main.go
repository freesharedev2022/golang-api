package main

import (
	"golang-api/controllers"
	"golang-api/middleware"
	"golang-api/models"
	"log"

	docs "golang-api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(middleware.XssMiddleware())
	if gin.Mode() == gin.ReleaseMode {
		r.Use(middleware.SecurityMiddleware())
	}
	r.Use(middleware.CorsMiddleware())

	models.ConnectDatabase()

	docs.SwaggerInfo.BasePath = "/"
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

	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
