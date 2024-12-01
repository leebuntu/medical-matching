package objects

import "medical-matching/constants/dto"

type Hospital struct {
	ID                 int
	Name               string
	OwnerName          string
	Address            string
	PostalCode         string
	Longitude          float64
	Latitude           float64
	ContactPhoneNumber string
	Symptoms           []*Symptom
	WaitingPerson      int
	ReviewStat         ReviewStat
	Facility           HospitalFacility
	OpenTime           []*OpenTime
}

type ReviewStat struct {
	AverageRating   float64
	TotalRating     int
	ReviewCount     int
	RatingStability float64
}

type HospitalFacility struct {
	HaveParkingLot int
}

type OpenTime struct {
	DayOfWeek int
	OpenTime  string
	CloseTime string
}

func (h *Hospital) CalculateReviewScore() float64 {
	return 0
}

func (h *Hospital) GetDTOHospitalDetail() dto.HospitalDetail {
	openTimes := make([]*dto.OpenTime, 0)
	for _, openTime := range h.OpenTime {
		openTimes = append(openTimes, &dto.OpenTime{
			DayOfWeek: openTime.DayOfWeek,
			OpenTime:  openTime.OpenTime,
			CloseTime: openTime.CloseTime,
		})
	}

	return dto.HospitalDetail{
		Name:               h.Name,
		OwnerName:          h.OwnerName,
		Address:            h.Address,
		ContactPhoneNumber: h.ContactPhoneNumber,
		WaitingPerson:      h.WaitingPerson,
		ReviewStat: &dto.ReviewStat{
			TotalAverageRating: float64(h.ReviewStat.TotalRating) / float64(h.ReviewStat.ReviewCount),
			Count:              h.ReviewStat.ReviewCount,
		},
		OpenTime: openTimes,
		Facility: &dto.HospitalFacility{
			HaveParkingLot: h.Facility.HaveParkingLot,
		},
	}
}

func (h *Hospital) GetDTOHospitalBrief() dto.HospitalBrief {
	return dto.HospitalBrief{
		Name:          h.Name,
		OwnerName:     h.OwnerName,
		Address:       h.Address,
		WaitingPerson: h.WaitingPerson,
	}
}
