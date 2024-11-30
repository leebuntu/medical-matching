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
		ctx.JSON(http.StatusOK, gin.H{"hospitals": hospitals})
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

		ctx.JSON(http.StatusOK, dto.HospitalDetail{
			Name:               hospital.Name,
			OwnerName:          hospital.OwnerName,
			Address:            hospital.Address,
			ContactPhoneNumber: hospital.ContactPhoneNumber,
			WaitingPerson:      hospital.WaitingPerson,
			//TODOOpenTime:           hospital.OpenTime,
		})
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

		ctx.JSON(http.StatusOK, dto.HospitalBrief{
			Name:          hospital.Name,
			OwnerName:     hospital.OwnerName,
			Address:       hospital.Address,
			WaitingPerson: hospital.WaitingPerson,
		})
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
