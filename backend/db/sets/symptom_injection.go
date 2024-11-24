package sets

import (
	"database/sql"
	"medical-matching/controller/hospital"
)

type SymptomInjection struct {
	db *sql.DB
}

func NewSymptomInjection(db *sql.DB) *SymptomInjection {
	return &SymptomInjection{db: db}
}

func (s *SymptomInjection) InjectSymptoms(hospitalID int, symptoms []hospital.Symptom) error {

}
