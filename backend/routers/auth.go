package routers

import (
	"MedicalMatching/constants"
	"MedicalMatching/constants/dto"
	"MedicalMatching/db"
	"MedicalMatching/db/auth"
	"MedicalMatching/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest dto.LoginRequest

		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": constants.BadRequest,
			})
			return
		}

		db, err := db.GetDBManager().GetDB(constants.UserDB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": constants.InternalServerError,
			})
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

		if err := c.ShouldBindJSON(&registerRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": constants.BadRequest,
			})
			return
		}

		db, err := db.GetDBManager().GetDB(constants.UserDB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": constants.InternalServerError,
			})
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
