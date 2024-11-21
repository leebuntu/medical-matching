package sets

import "database/sql"

type HospitalInjection struct {
	db *sql.DB
}

func NewHospitalInjection(db *sql.DB) *HospitalInjection {
	return &HospitalInjection{db: db}
}

func (h *HospitalInjection) InjectTestHospital() error {
	return nil
}
