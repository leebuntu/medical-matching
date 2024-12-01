package objects

type Symptom struct {
	ID   int
	Name string
}

type WeightedScore struct {
	TotalScore  float64
	ContentRank []int
}

type ScoredHospital struct {
	HospitalID    int
	Score         float64
	Content       []int
	WaitingPerson int
}
