package dto

type Symptoms struct {
	KnownSymptoms      []int  `json:"known_symptoms"`
	AdditionalSymptoms string `json:"additional_symptoms"`
}

type MatchingRequest struct {
	BasisLongitude float64  `json:"basis_longitude"`
	BasisLatitude  float64  `json:"basis_latitude"`
	Radius         float64  `json:"radius"`
	Symptoms       Symptoms `json:"symptoms"`
}

type PoolingResponseNotCompleted struct {
	State int `json:"state"`
}

type PoolingResponseCompleted struct {
	State         int      `json:"state"`
	HospitalID    int      `json:"hospital_id"`
	ContentOption []string `json:"content_option"`
	WaitingPerson int      `json:"waiting_person"`
}
