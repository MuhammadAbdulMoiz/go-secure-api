package routes

import (
	"go-secure-api/controllers"
	"go-secure-api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())
	router.GET("/user", controllers.GetUsers())
	router.GET("/user/:userID", controllers.GetUser())
	router.POST("/user/logout", controllers.Logout())
}
