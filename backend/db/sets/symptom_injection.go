package sets

import (
	"database/sql"
	"medical-matching/constants/objects"
)

type SymptomInjection struct {
	db *sql.DB
}

func NewSymptomInjection(db *sql.DB) *SymptomInjection {
	return &SymptomInjection{db: db}
}

func (s *SymptomInjection) InjectSymptoms(hospitalID int, symptoms []objects.Symptom) error {

}
