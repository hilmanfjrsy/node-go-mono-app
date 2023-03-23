package main

import (
	"fetch-app/controllers"
	"fetch-app/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()
	r := gin.Default()
	resources := r.Group("/resources")

	resources.GET("", middleware.AuthMiddleware, controllers.GetResource)
	resources.GET("/aggregate", middleware.AuthMiddleware, middleware.OnlyAdmin, controllers.GetResourceAggregate)
	r.Run()
}
