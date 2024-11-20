package routers

import (
	"MedicalMatching/constants"
	"MedicalMatching/constants/dto"
	"MedicalMatching/db"
	"MedicalMatching/db/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("userID")

		db, err := db.GetDBManager().GetDB(constants.UserDB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		userService := user.NewUserService(db)
		profile, err := userService.GetUserProfile(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		c.JSON(http.StatusOK, profile)

	}
}

func UpdateUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("userID")

		var profile dto.UserProfileUpdate
		if err := c.ShouldBindJSON(&profile); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return
		}

		db, err := db.GetDBManager().GetDB(constants.UserDB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		userService := user.NewUserService(db)
		err = userService.UpdateUserProfile(userID, profile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": constants.UpdateProfileSuccess})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DeleteUserProfile",
		})
	}
}

func AddPaymentMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GetPaymentMethodList() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func DeletePaymentMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
