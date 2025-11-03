package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType != "ADMIN" && uid != userId {
		return errors.New("unauthorized access")
	}

	err = CheckUserType(c, userType)
	return err

}

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil

	if userType != role {
		return errors.New("unauthorized access")
	}

	return err
}
