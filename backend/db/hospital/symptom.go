package hospital

import (
	"database/sql"
	"medical-matching/constants"
	"medical-matching/constants/objects"
	"medical-matching/db"
	"sync"
)

type SymptomManager struct {
	db       *sql.DB
	symptoms map[int]string
}

var symptomOnce sync.Once
var symptomInstance *SymptomManager

func GetSymptomManager() *SymptomManager {
	symptomOnce.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.HospitalDB)
		if err != nil {
			return
		}
		symptomInstance = &SymptomManager{db: db}
	})
	return symptomInstance
}

func (m *SymptomManager) GetSymptom(id int) *objects.Symptom {
	return &objects.Symptom{ID: id, Name: m.symptoms[id]}
}

func (m *SymptomManager) ResetSymptomManager() error {
	rows, err := m.db.Query("SELECT * FROM symptom")
	if err != nil {
		return err
	}

	defer rows.Close()

	symptoms := make(map[int]string)

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return err
		}
		symptoms[id] = name
	}

	return nil
}
