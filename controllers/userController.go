package controllers

import (
	"go-secure-api/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.DBClient, "users")
var validate = validator.New()

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func hashPassword(password string) string {
	return password
}

func verifyPassword(userPassword string, providedPassword string) bool {
	return true
}
