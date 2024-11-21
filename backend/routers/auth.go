package routers

import (
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/db/auth"
	"medical-matching/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest dto.LoginRequest

		db, _, err := initRouteHandler(c, &loginRequest, constants.UserDB)
		if err != nil {
			return
		}

		loginService := auth.NewAuthService(db)

		userID, err := loginService.Login(&loginRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": constants.WrongAccountOrPassword,
			})
			return
		}

		token, err := middlewares.GenerateJWT(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": constants.InternalServerError,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var registerRequest dto.RegisterRequest

		db, _, err := initRouteHandler(c, &registerRequest, constants.UserDB)
		if err != nil {
			return
		}

		registerService := auth.NewAuthService(db)

		registerResult, err := registerService.Register(&registerRequest)
		if !registerResult {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": constants.DuplicateUser,
			})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": constants.InternalServerError,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": constants.WelcomeRegister,
		})
	}
}
