package medical

type MatchingRequest struct {
	BasisLocation   string   `json:"basis_location"`
	Symptoms        string   `json:"symptoms"`
	HospitalOptions []string `json:"hospital_options"`
}
