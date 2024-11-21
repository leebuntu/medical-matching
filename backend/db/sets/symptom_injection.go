package sets

import (
	"MedicalMatching/constants/objects"
	"database/sql"
)

type SymptomInjection struct {
	db *sql.DB
}

func NewSymptomInjection(db *sql.DB) *SymptomInjection {
	return &SymptomInjection{db: db}
}

func (s *SymptomInjection) InjectSymptoms(hospitalID int, symptoms []objects.Symptom) error {

}
