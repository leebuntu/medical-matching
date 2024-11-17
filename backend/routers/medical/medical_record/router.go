package medical_record

import "github.com/gin-gonic/gin"

func GetMedicalRecordHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		recordID := c.Param("recordID")

		if recordID == "" {

		} else {

		}

	}
}
