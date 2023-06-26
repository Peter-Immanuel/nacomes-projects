package handlers

import "github.com/gin-gonic/gin"

func APIHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Backend is healthty",
	})
}
