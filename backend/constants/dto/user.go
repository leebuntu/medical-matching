package dto

type UserProfile struct {
	Username       string `json:"username"`
	ProfileURL     string `json:"profile_url"`
	PhoneNumber    string `json:"phone_number"`
	HomeAddress    string `json:"home_address"`
	Email          string `json:"email"`
	CardID         string `json:"card_id"`
	Candy          int    `json:"candy"`
	PriorityOption []int  `json:"priority_option"`
}

type UserProfileUpdate struct {
	ProfileURL     string `json:"profile_url"`
	PhoneNumber    string `json:"phone_number"`
	HomeAddress    string `json:"home_address"`
	PostalCode     string `json:"postal_code"`
	CardID         string `json:"card_id"`
	PriorityOption []int  `json:"priority_option"`
}

type PaymentMethod struct {
	CardHolderName string `json:"card_holder_name"`
	CardNumber     string `json:"card_number"`
	ExpDate        string `json:"exp_date"`
	Cvv            string `json:"cvv"`
}

type RetrievePaymentMethod struct {
	CardID     string `json:"card_id"`
	CardNumber string `json:"card_number"`
	ExpDate    string `json:"exp_date"`
}
