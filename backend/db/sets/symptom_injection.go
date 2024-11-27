package sets

import (
	"database/sql"
)

type SymptomInjection struct {
	db *sql.DB
}

var symptoms = []string{
	"headache", "fever", "whirl", "lump", "hair_loss",
}

func NewSymptomInjection(db *sql.DB) *SymptomInjection {
	return &SymptomInjection{db: db}
}

func (s *SymptomInjection) alreadyInjected() (bool, error) {
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM symptom").Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *SymptomInjection) InjectSymptoms() error {
	alreadyInjected, err := s.alreadyInjected()
	if err != nil {
		return err
	}
	if alreadyInjected {
		return nil
	}

	for _, symptom := range symptoms {
		_, err := s.db.Exec("INSERT INTO symptom (name) VALUES (?)", symptom)
		if err != nil {
			return err
		}
	}

	return nil
}
