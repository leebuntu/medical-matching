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
		rq := dto.MatchingRequest{}
		userID, err := CheckBindData(ctx, &rq)
		if err != nil {
			return
		}

		mm := matching.GetMatchingManager()
		m, err := mm.CreateMatching(userID, &rq)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return
		}

		go m.StartMatching()

		ctx.JSON(http.StatusOK, gin.H{"matching_id": m.GetMatchingID()})
	}
}

func GetMatching() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		matchingID := ctx.Param("matchingID")

		m := matching.GetMatchingManager()
		matching, err := m.GetMatching(matchingID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": constants.NotFound})
			return
		}

		if matching.GetState() == constants.MatchingCompleted {
			ctx.JSON(http.StatusOK, matching.GetCompleteResult())
		} else if matching.GetState() == constants.MatchingFailed {
			ctx.JSON(http.StatusOK, dto.PoolingResponseNotCompleted{State: constants.MatchingFailed})
			m.RemoveMatching(matchingID)
		} else {
			ctx.JSON(http.StatusOK, dto.PoolingResponseNotCompleted{State: 0})
		}
	}
}
