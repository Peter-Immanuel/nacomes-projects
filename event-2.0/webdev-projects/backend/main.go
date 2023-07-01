package main

import (
	"log"

	"github.com/Peter-Immanuel/nacomes-projects/event-2.0/webdev-projects/backend/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.Use(cors.Default())
	router.GET("/", handlers.APIHealth)
	router.POST("/recommendation", handlers.GetRecommendation)
	router.POST("/intruder", handlers.IntruderAlertHandler)
	router.Run()
}
