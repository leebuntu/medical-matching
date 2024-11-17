package auth

import (
	"MedicalMatching/constants"
	"MedicalMatching/constants/dto/auth"
	"MedicalMatching/db"
	dbauth "MedicalMatching/db/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var registerRequest auth.RegisterRequest

		if err := c.BindJSON(&registerRequest); err != nil {
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

		registerService := dbauth.NewAuthService(db)

		registerResult, err := registerService.Register(&registerRequest)
		if !registerResult {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": constants.DuplicateUser,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": constants.WelcomeRegister,
		})
	}
}
