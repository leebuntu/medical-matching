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
		db, userID, err := initRouteHandler(c, nil, constants.UserDB)
		if err != nil {
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
		profile := dto.UserProfileUpdate{}

		db, userID, err := initRouteHandler(c, &profile, constants.UserDB)
		if err != nil {
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
		var card dto.PaymentMethod

		db, userID, err := initRouteHandler(c, &card, constants.UserDB)

		paymentService := user.NewPaymentService(db)
		err = paymentService.AddPaymentMethod(userID, card)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": constants.AddPaymentMethodSuccess})
	}
}

func GetPaymentMethodList() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, userID, err := initRouteHandler(c, nil, constants.UserDB)
		if err != nil {
			return
		}

		paymentService := user.NewPaymentService(db)
		paymentMethods, err := paymentService.GetPaymentMethodList(userID)
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

		db, err := db.GetDBManager().GetDB(constants.UserDB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		paymentService := user.NewPaymentService(db)
		err = paymentService.DeletePaymentMethod(userID, cardID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": constants.DeletePaymentMethodSuccess})
	}
}
