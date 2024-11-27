package hospital

import (
	"database/sql"
	"medical-matching/constants"
	"medical-matching/constants/objects"
	"medical-matching/controller/matching"
	"medical-matching/db"
	"sync"
)

type HospitalManager struct {
	db        *sql.DB
	hospitals map[int]*objects.Hospital
}

var once sync.Once
var hospitalInstance *HospitalManager

func GetHospitalManager() *HospitalManager {
	once.Do(func() {
		hospitalInstance = &HospitalManager{hospitals: make(map[int]*objects.Hospital)}
	})
	return hospitalInstance
}

func (m *HospitalManager) getHospital() error {
	rows, err := m.db.Query("SELECT * FROM hospital")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var hospital objects.Hospital
		err := rows.Scan(&hospital.ID, &hospital.Name, &hospital.OwnerName, &hospital.Address, &hospital.PostalCode, &hospital.ContactPhone)
		if err != nil {
			return err
		}
		m.hospitals[hospital.ID] = &hospital
	}

	return nil
}

func (m *HospitalManager) InitHospitalManager() error {
	db, err := db.GetDBManager().GetDB(constants.HospitalDB)
	m.db = db

	if err != nil {
		return err
	}

	m.getHospital()

	return nil
}

func (m *HospitalManager) FilteringHospital(composer matching.Composer) *objects.Hospital {
	return nil
}
