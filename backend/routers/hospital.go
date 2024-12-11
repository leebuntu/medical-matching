package routers

import (
	"medical-matching/constants"
	"medical-matching/constants/dto"
	ch "medical-matching/controller/hospital"
	"medical-matching/db/providers"
	"medical-matching/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetHospitalList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		longitude, latitude, radius, err := utils.ParseGEOQuery(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return
		}

		hm := ch.GetHospitalManager()
		hospitals, err := hm.GetHospitals(longitude, latitude, radius)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		dtoHospitals := make([]*dto.HospitalDetail, 0)
		for _, hospital := range hospitals {
			dtoHospitals = append(dtoHospitals, hospital.GetDTOHospitalDetail())
		}

		ctx.JSON(http.StatusOK, dto.HospitalListResponse{Count: len(hospitals), Hospitals: dtoHospitals})
	}
}

func GetHospitalDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hospitalID, err := strconv.Atoi(ctx.Param("hospitalID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return
		}

		hm := ch.GetHospitalManager()
		hospital := hm.GetHospital(hospitalID)
		if hospital == nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": constants.NotFound})
			return
		}

		ctx.JSON(http.StatusOK, hospital.GetDTOHospitalDetail())
	}
}

func GetBriefHospital() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hospitalID, err := strconv.Atoi(ctx.Param("hospitalID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return
		}

		hm := ch.GetHospitalManager()
		hospital := hm.GetHospital(hospitalID)
		if hospital == nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": constants.NotFound})
			return
		}

		ctx.JSON(http.StatusOK, hospital.GetDTOHospitalBrief())
	}
}

func GetHospitalReview() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hospitalID := ctx.Param("hospitalID")
		if hospitalID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return
		}

		page := ctx.Query("page")
		if page == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
		}
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return
		}

		rm := providers.GetReviewProvider()
		reviews, err := rm.GetReview(hospitalID, pageInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.InternalServerError})
			return
		}

		ctx.JSON(http.StatusOK, dto.ReviewResponse{
			Count:       len(reviews),
			CurrentPage: pageInt,
			Reviews:     reviews,
		})
	}
}
