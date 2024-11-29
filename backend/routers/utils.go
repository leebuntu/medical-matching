package routers

import (
	"medical-matching/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckBindData(c *gin.Context, bindData interface{}) (int, error) {
	userID := c.GetInt("userID")

	if bindData != nil {
		if err := c.ShouldBindJSON(bindData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return 0, err
		}
	}

	return userID, nil
}
