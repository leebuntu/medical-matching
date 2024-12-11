package routers

import (
	"fmt"
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/controller/hospital"
	"medical-matching/controller/matching"
	"medical-matching/db/providers"
	"medical-matching/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateMatching() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rq := dto.MatchingRequest{}
		userID, err := utils.CheckBindData(ctx, &rq)
		if err != nil {
			return
		}

		mm := matching.GetMatchingManager()
		m, err := mm.CreateMatching(userID, rq.Symptoms.KnownSymptoms)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return
		}

		hospitals, err := hospital.GetHospitalManager().GetHospitals(rq.BasisLongitude, rq.BasisLatitude, rq.Radius)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		up := providers.GetUserProvider()
		priority, err := up.GetPriorityByID(userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		go func() {
			m.StartMatching(rq.BasisLongitude, rq.BasisLatitude, priority, hospitals)
			rp := providers.GetRecordProvider()

			symptoms := rq.Symptoms.KnownSymptoms
			var symptomLists string
			for _, symptom := range symptoms {
				symptomLists += strconv.Itoa(symptom) + ", "
			}
			symptomLists = strings.TrimSuffix(symptomLists, ", ")

			err = rp.AddRecord(userID, m.GetCompleteResult().HospitalID, symptomLists)
			if err != nil {
				fmt.Println(err)
			}
		}()

		ctx.JSON(http.StatusOK, gin.H{"matching_id": m.GetMatchingID()})
	}
}

func GetMatching() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		matchingID := ctx.Param("matchingID")
		userID := ctx.GetInt("userID")

		m := matching.GetMatchingManager()
		matching, err := m.GetMatching(matchingID, userID)
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

func GetAllMatching() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetInt("userID")

		mm := matching.GetMatchingManager()
		matchings := mm.GetAllMatching(userID)

		ctx.JSON(http.StatusOK, dto.MatchingListResponse{Count: len(matchings), MatchingIDs: matchings})
	}
}

func EndMatching() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		matchingID := ctx.Param("matchingID")

		mm := matching.GetMatchingManager()
		err := mm.RemoveMatching(matchingID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": constants.NotFound})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": constants.EndMatchingSuccess})
	}
}
