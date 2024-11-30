package objects

type Hospital struct {
	ID                 int
	Name               string
	OwnerName          string
	Address            string
	PostalCode         string
	Longitude          float64
	Latitude           float64
	ContactPhoneNumber string
	SymptomIDs         []int
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

type ScoredHospital struct {
	HospitalID    int
	Score         float64
	Content       []string
	WaitingPerson int
}

type OpenTime struct {
	DayOfWeek int
	OpenTime  string
	CloseTime string
}
