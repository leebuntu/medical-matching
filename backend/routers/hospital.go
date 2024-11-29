package routers

import (
	"medical-matching/constants"
	"medical-matching/db/hospital"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHospitalList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// NOTE: currently show all hospitals
		hm := hospital.GetHospitalManager()
		hospitals, err := hm.GetHospitals()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"hospitals": hospitals})
	}
}

func GetHospitalDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//hospitalID := ctx.Param("hospitalID")
		// TODO: get hospital detail
	}
}
