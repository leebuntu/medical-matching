package dto

type Symptoms struct {
	KnownSymptoms      []int  `json:"known_symptoms" binding:"required"`
	AdditionalSymptoms string `json:"additional_symptoms"`
}

type MatchingRequest struct {
	BasisLongitude float64  `json:"basis_longitude" binding:"required"`
	BasisLatitude  float64  `json:"basis_latitude" binding:"required"`
	Radius         float64  `json:"radius" binding:"required"`
	Symptoms       Symptoms `json:"symptoms" binding:"required"`
}

type MatchingListResponse struct {
	Count       int      `json:"count"`
	MatchingIDs []string `json:"matching_ids"`
}

type PoolingResponseNotCompleted struct {
	State int `json:"state"`
}

type PoolingResponseCompleted struct {
	State         int   `json:"state"`
	HospitalID    int   `json:"hospital_id"`
	ContentOption []int `json:"content_option"`
	WaitingPerson int   `json:"waiting_person"`
}
