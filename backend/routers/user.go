package routers

import (
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/db/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("userID")

		userService := user.GetService()
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
		profile := dto.UserProfileUpdate{}

		userID, err := CheckBindData(c, &profile)
		if err != nil {
			return
		}

		userService := user.GetService()
		err = userService.UpdateUserProfile(userID, &profile)

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
			"message": "Not implemented",
		})
	}
}

func AddPaymentMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		var card dto.PaymentMethod

		userID, err := CheckBindData(c, &card)
		if err != nil {
			return
		}

		userService := user.GetService()
		err = userService.AddPaymentMethod(userID, &card)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": constants.AddPaymentMethodSuccess})
	}
}

func GetPaymentMethodList() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("userID")

		userService := user.GetService()
		paymentMethods, err := userService.GetPaymentMethodList(userID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		c.JSON(http.StatusOK, paymentMethods)
	}
}

func DeletePaymentMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("userID")
		cardID := c.Param("cardID")

		userService := user.GetService()
		err := userService.DeletePaymentMethod(userID, cardID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": constants.DeletePaymentMethodSuccess})
	}
}
