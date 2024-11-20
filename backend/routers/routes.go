package routers

import (
	"MedicalMatching/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/login", LoginHandler())
			authGroup.POST("/register", RegisterHandler())
			//auth.POST("/logout", LogoutHandler)
		}

		userGroup := v1.Group("/users")
		{
			userGroup.Use(middlewares.AuthMiddleware())
			userGroup.GET("/me", GetUserProfile())
			userGroup.PUT("/me", UpdateUserProfile())
			//userGroup.DELETE("/me", user.DeleteUserProfile())
		}

		paymentGroup := userGroup.Group("/payment-methods")
		{
			paymentGroup.Use(middlewares.AuthMiddleware())
			paymentGroup.POST("/", AddPaymentMethod())
			paymentGroup.GET("/", GetPaymentMethodList())
			paymentGroup.DELETE("/", DeletePaymentMethod())
		}

	}
}
