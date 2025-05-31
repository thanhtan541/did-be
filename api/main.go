package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load configuration: %v", err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // Default port is :8080
}
