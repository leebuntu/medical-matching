package auth

type LoginRequest struct {
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

type RegisterRequest struct {
	Email          string `json:"email_address"`
	HashedPassword string `json:"hashed_password"`
	Username       string `json:"username"`
	PhoneNumber    string `json:"phone_number"`
	HomeAddress    string `json:"home_address"`
	PostalCode     string `json:"postal_code"`
}
