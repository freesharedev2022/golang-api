package main

import (
	"golang-api/middleware"
	"golang-api/models"
	"golang-api/routes"
	"log"

	docs "golang-api/docs"

	"github.com/gin-gonic/gin"
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
	// config routes
	routes.Route(r)

	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
