package dto

type UserProfile struct {
	Username       string `json:"username"`
	ProfileURL     string `json:"profile_url"`
	PhoneNumber    string `json:"phone_number"`
	HomeAddress    string `json:"home_address"`
	Email          string `json:"email"`
	PaymentMethod  string `json:"payment_method"`
	Candy          int    `json:"candy"`
	PriorityOption []int  `json:"priority_option"`
}

type UserProfileUpdate struct {
	ProfileURL     string `json:"profile_url"`
	PhoneNumber    string `json:"phone_number"`
	HomeAddress    string `json:"home_address"`
	PaymentID      string `json:"payment_id"`
	PriorityOption []int  `json:"priority_option"`
}

type PaymentMethod struct {
	CardHolderName string `json:"card_holder_name"`
	CardNumber     string `json:"card_number"`
	ExpDate        string `json:"exp_date"`
	Cvv            string `json:"cvv"`
}

type DeletePaymentMethod struct {
	PaymentID string `json:"payment_id"`
}

type PaymentMethodList struct {
	PaymentMethods []PaymentMethod `json:"payment_methods"`
}
