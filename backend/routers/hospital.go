package routers

import "github.com/gin-gonic/gin"

func GetHospitalList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: get hospital list
	}
}

func GetHospitalDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//hospitalID := ctx.Param("hospitalID")
		// TODO: get hospital detail
	}
}
