package main

import (
	"github.com/Peter-Immanuel/nacomes-projects/event-2.0/webdev-projects/backend/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.GET("/", handlers.APIHealth)
	router.POST("/recommendation", handlers.GetRecommendation)
	//router.Run(":3001")
	router.Run()
}
