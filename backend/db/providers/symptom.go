package providers

import (
	"database/sql"
	"medical-matching/constants"
	"medical-matching/db"
	"medical-matching/objects"
	"sync"
)

type SymptomProvider struct {
	db *sql.DB
}

var symptomOnce sync.Once
var symptomInstance *SymptomProvider

func GetSymptomProvider() *SymptomProvider {
	symptomOnce.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.HospitalDB)
		if err != nil {
			return
		}
		symptomInstance = &SymptomProvider{db: db}
	})
	return symptomInstance
}

func (p *SymptomProvider) FetchSymptoms() ([]*objects.Symptom, error) {
	rows, err := p.db.Query("SELECT * FROM symptom")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	symptoms := make([]*objects.Symptom, 0)

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		symptoms = append(symptoms, &objects.Symptom{ID: id, Name: name})
	}

	return symptoms, nil
}
