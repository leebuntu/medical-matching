package objects

type Hospital struct {
	ID            int
	Name          string
	OwnerName     string
	Address       string
	PostalCode    string
	ContactPhone  string
	Symptoms      []int
	WaitingPerson int
	ReviewStat    ReviewStat
	Facility      HospitalFacility
	Reserved      int
	Reserved2     int
	Reserved3     int
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
