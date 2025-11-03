package routes

import (
	"go-secure-api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("user/register", controllers.Register())
	router.POST("/user/login", controllers.Login())
}
