package dto

type Symptoms struct {
	KnownSymptoms      []int  `json:"known_symptoms"`
	AdditionalSymptoms string `json:"additional_symptoms"`
}

type MatchingRequest struct {
	BasisLocation string   `json:"basis_location"`
	Symptoms      Symptoms `json:"symptoms"`
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
