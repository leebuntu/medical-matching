package routers

import (
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/controller/matching"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMatching() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetInt("userID")

		rq := dto.MatchingRequest{}
		if err := ctx.ShouldBindJSON(&rq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return
		}

		m := matching.GetMatchingManager()
		m.CreateMatching(userID, &rq)
		m.StartMatching(userID)
	}
}

func GetMatching() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetInt("userID")
		matchingID := ctx.Param("matchingID")

		m := matching.GetMatchingManager()
		matching, err := m.GetMatching(userID, matchingID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": constants.NotFound})
			return
		}

		if matching.GetState() == constants.MatchingCompleted {
			ctx.JSON(http.StatusOK, dto.PoolingResponseCompleted{State: constants.MatchingCompleted})
		} else {
			ctx.JSON(http.StatusOK, dto.PoolingResponseNotCompleted{State: 0})
		}
	}
}
