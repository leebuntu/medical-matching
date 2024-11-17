package auth

import (
	"MedicalMatching/constants"
	"MedicalMatching/constants/dto/auth"
	"MedicalMatching/db"
	dbauth "MedicalMatching/db/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest auth.LoginRequest

		if err := c.BindJSON(&loginRequest); err != nil {
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

		loginService := dbauth.NewAuthService(db)

		userID, err := loginService.Login(&loginRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": constants.WrongAccountOrPassword,
			})
			return
		}

		token, err := GenerateJWT(userID)
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
