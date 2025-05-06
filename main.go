package main

import (
	"haikal/backend-api/config"
	"haikal/backend-api/database"

	"github.com/gin-gonic/gin"
)

func main() {

	// load .env
	config.LoadEnv()

	// inisialisasi database
	database.InitDB()

	// inisialisasi gin
	router := gin.Default()

	// membuat route dengan method get
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.Run(":" + config.GetEnv("APP_PORT", "3300"))
}
