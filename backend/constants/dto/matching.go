package dto

const (
	Waiting = iota + 1
	Distance
	Review
	HaveParkingLot
	LeastWalk
)

var Weights = map[int]float64{
	1: 3.0,
	2: 2.0,
	3: 1.5,
}

const (
	TotalPriority = 5
)

const (
	PerWatingPersonScore = -5
	HaveParkingLotScore  = 30
	PerWalkMinuteScore   = -5
)

const (
	BeforeMatching = iota
	StartMatching
	Reserved
	MatchingCompleted
)

type Symptoms struct {
	KnownSymptoms      []string `json:"known_symptoms"`
	AdditionalSymptoms string   `json:"additional_symptoms"`
}

type MatchingRequest struct {
	BasisLocation string   `json:"basis_location"`
	Symptoms      Symptoms `json:"symptoms"`
}

type PoolingResponse struct {
	PoolingID string `json:"pooling_id"`
}

type PoolingRequest struct {
	PoolingID string `json:"pooling_id"`
}
