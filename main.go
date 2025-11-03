package main

import (
	routes "go-secure-api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api/v1/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Secure API v1"})
	})

	router.GET("api/v2/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Secure API v2"})
	})
	router.Run(":" + port)

}
