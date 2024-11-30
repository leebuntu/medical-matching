package dto

type HospitalDetail struct {
	Name               string            `json:"name"`
	OwnerName          string            `json:"owner_name"`
	Address            string            `json:"address"`
	ContactPhoneNumber string            `json:"contact_phone_number"`
	WaitingPerson      int               `json:"waiting_person"`
	OpenTime           []*OpenTime       `json:"open_time"`
	Facility           *HospitalFacility `json:"facility"`
}

type HospitalBrief struct {
	Name          string `json:"name"`
	OwnerName     string `json:"owner_name"`
	Address       string `json:"address"`
	WaitingPerson int    `json:"waiting_person"`
}

type ReviewResponse struct {
	Count       int           `json:"count"`
	CurrentPage int           `json:"current_page"`
	Reviews     []*ReviewStat `json:"reviews"`
}

type ReviewStat struct {
	Rating        int      `json:"rating"`
	VisitedDate   string   `json:"visited_date"`
	ProfileURL    string   `json:"profile_url"`
	ProfileName   string   `json:"profile_name"`
	ReviewContext string   `json:"review_context"`
	ReviewPhoto   []string `json:"review_photo"`
}

type OpenTime struct {
	DayOfWeek int    `json:"day_of_week"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}

type HospitalFacility struct {
	HaveParkingLot int `json:"have_parking_lot"`
}
