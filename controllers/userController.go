package controllers

import (
	"context"
	"go-secure-api/database"
	"go-secure-api/helpers"
	"go-secure-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
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
	return func(c *gin.Context) {
		userId := c.Param("userID")
		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error(err)})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"userID": userId}).Decode(&user)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error(err)})
			return
		}
		c.JSON(http.StatusOK, user)

	}
}

func hashPassword(password string) string {
	return password
}

func verifyPassword(userPassword string, providedPassword string) bool {
	return true
}
