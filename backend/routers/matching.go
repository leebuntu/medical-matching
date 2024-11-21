package routers

import "github.com/gin-gonic/gin"

func CreateMatching() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetInt("userID")
	}
}

func GetMatching() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetInt("userID")
		matchingID := ctx.Param("matchingID")
	}
}
