package dto

import "medical-matching/constants/objects"

type HospitalDetail struct {
	Name               string              `json:"name"`
	OwnerName          string              `json:"owner_name"`
	Address            string              `json:"address"`
	ContactPhoneNumber string              `json:"contact_phone_number"`
	WaitingPerson      int                 `json:"waiting_person"`
	OpenTime           []*objects.OpenTime `json:"open_time"`
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
