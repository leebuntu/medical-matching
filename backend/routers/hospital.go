package routers

import (
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/db/hospital"
	"net/http"
	"strconv"

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
		hospitalID, err := strconv.Atoi(ctx.Param("hospitalID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return
		}

		hm := hospital.GetHospitalManager()
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
			OpenTime:           hospital.OpenTime,
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

		hm := hospital.GetHospitalManager()
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

		rm := hospital.GetReviewManager()
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
