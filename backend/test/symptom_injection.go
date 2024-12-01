package test

import (
	"database/sql"
	"encoding/csv"
	"medical-matching/constants"
	"medical-matching/db"
	"os"
	"sync"
)

type SymptomInjection struct {
	db *sql.DB
}

var symptomOnce sync.Once
var symptomInstance *SymptomInjection

func GetSymptomInjection() *SymptomInjection {
	symptomOnce.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.HospitalDB)
		if err != nil {
			return
		}
		symptomInstance = &SymptomInjection{db: db}
	})
	return symptomInstance
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

	file, err := os.Open(constants.TestDataPath + constants.SymptomTestData)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		return err
	}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		_, err = s.db.Exec("INSERT INTO symptom (name) VALUES (?)", record[0])
		if err != nil {
			return err
		}
	}

	return nil
}
