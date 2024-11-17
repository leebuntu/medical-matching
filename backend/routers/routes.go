package routers

import (
	"MedicalMatching/routers/auth"
	"MedicalMatching/routers/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/login", auth.LoginHandler())
			authGroup.POST("/register", auth.RegisterHandler())
			//auth.POST("/logout", LogoutHandler)
		}

		userGroup := v1.Group("/users")
		{
			userGroup.Use(auth.AuthMiddleware())
			userGroup.GET("/me", user.GetUserProfile())
			userGroup.PUT("/me", user.UpdateUserProfile())
			userGroup.DELETE("/me", user.DeleteUserProfile())
		}

	}
}
