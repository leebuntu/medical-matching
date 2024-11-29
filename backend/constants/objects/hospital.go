package objects

type Hospital struct {
	ID                 int
	Name               string
	OwnerName          string
	Address            string
	PostalCode         string
	ContactPhoneNumber string
	Symptoms           []*Symptom
	WaitingPerson      int
	ReviewStat         ReviewStat
	Facility           HospitalFacility
	OpenTime           []OpenTime
}

type ReviewStat struct {
	AverageRating float64
	TotalRating   int
	ReviewCount   int
	DXRating      float64
}

type HospitalFacility struct {
	HaveParkingLot bool
}

type Symptom struct {
	ID   int
	Name string
}

type ScoredHospital struct {
	HospitalID int
	Score      float64
	Content    []string
}

type OpenTime struct {
	DayOfWeek int    `json:"day_of_week"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}
