package routers

import (
	"MedicalMatching/constants"
	"MedicalMatching/db"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRouteHandler(c *gin.Context, bindData interface{}, dbName string) (*sql.DB, int, error) {
	userID := c.GetInt("userID")

	if bindData != nil {
		if err := c.ShouldBindJSON(bindData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return nil, 0, err
		}
	}

	db, err := db.GetDBManager().GetDB(dbName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
		return nil, 0, err
	}

	return db, userID, nil
}
