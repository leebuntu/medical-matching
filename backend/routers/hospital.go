package routers

import "github.com/gin-gonic/gin"

func GetHospitalList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetInt("userID")
	}
}
