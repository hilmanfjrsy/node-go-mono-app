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

	resources.Use(middleware.AuthMiddleware)
	resources.GET("", controllers.GetResource)
	resources.GET("/currency", controllers.GetResourceWithCurrency)
	resources.GET("/aggregate", middleware.OnlyAdmin, controllers.GetResourceAggregate)
	r.Run()
}
