package hospital

import (
	"MedicalMatching/constants"
	"MedicalMatching/db"
	"database/sql"
)

type Symptom struct {
	ID   int
	Name string
}

type SymptomManager struct {
	db       *sql.DB
	Symptoms map[int]Symptom
}

var symptomInstance *SymptomManager

func GetSymptomManager() *SymptomManager {
	once.Do(func() {
		symptomInstance = &SymptomManager{Symptoms: make(map[int]Symptom)}
	})
	return symptomInstance
}

func (m *SymptomManager) getSymptoms() error {
	rows, err := m.db.Query("SELECT * FROM symptom")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var symptom Symptom
		err := rows.Scan(&symptom.ID, &symptom.Name)
		if err != nil {
			return err
		}
		m.Symptoms[symptom.ID] = symptom
	}

	return nil
}

func (m *SymptomManager) GetSymptom(id int) Symptom {
	return m.Symptoms[id]
}

func (m *SymptomManager) InitSymptomManager() error {
	db, err := db.GetDBManager().GetDB(constants.HospitalDB)
	m.db = db
	if err != nil {
		return err
	}

	err = m.getSymptoms()
	if err != nil {
		return err
	}

	return nil
}
