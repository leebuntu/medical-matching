package routers

import (
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/db/providers"
	"medical-matching/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("userID")

		userService := providers.GetUserProvider()
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

		userID, err := utils.CheckBindData(c, &profile)
		if err != nil {
			return
		}

		userProvider := providers.GetUserProvider()
		err = userProvider.UpdateUserProfile(userID, &profile)

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

		userID, err := utils.CheckBindData(c, &card)
		if err != nil {
			return
		}

		userService := providers.GetUserProvider()
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

		userService := providers.GetUserProvider()
		paymentMethods, err := userService.GetPaymentMethodList(userID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		c.JSON(http.StatusOK, gin.H{"payment_methods": paymentMethods})
	}
}

func DeletePaymentMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("userID")
		cardID := c.Param("cardID")

		userService := providers.GetUserProvider()
		err := userService.DeletePaymentMethod(userID, cardID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": constants.DeletePaymentMethodSuccess})
	}
}
