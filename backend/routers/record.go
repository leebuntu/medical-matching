package routers

import (
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/db/providers"
	"medical-matching/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRecordList() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("userID")
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return
		}

		records, err := providers.GetRecordProvider().GetRecordList(userID, page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}
		c.JSON(http.StatusOK, dto.RecordListResponse{Count: len(records), CurrentPage: page, Records: records})
	}
}

func UpdateRecordNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &dto.UpdateRecordNotesRequest{}
		userID, err := utils.CheckBindData(c, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return
		}

		recordID, err := strconv.Atoi(c.Param("recordID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return
		}

		err = providers.GetRecordProvider().UpdateRecordNotes(recordID, userID, req.Notes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
