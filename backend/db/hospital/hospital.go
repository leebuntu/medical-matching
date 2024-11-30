package hospital

import (
	"database/sql"
	"errors"
	"medical-matching/constants"
	"medical-matching/constants/objects"
	"medical-matching/db"
	"sync"
)

type HospitalManager struct {
	db        *sql.DB
	hospitals map[int]*objects.Hospital
}

var hospitalOnce sync.Once
var hospitalInstance *HospitalManager

func GetHospitalManager() *HospitalManager {
	hospitalOnce.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.HospitalDB)
		if err != nil {
			return
		}
		hospitalInstance = &HospitalManager{db: db}
	})
	return hospitalInstance
}

func (m *HospitalManager) getHospitalBasicInfo() error {
	rows, err := m.db.Query("SELECT id, name, owner_name, address, postal_code, contact_phone_number FROM hospital")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		hospital := &objects.Hospital{}
		err := rows.Scan(&hospital.ID, &hospital.Name, &hospital.OwnerName, &hospital.Address, &hospital.PostalCode, &hospital.ContactPhoneNumber)
		if err != nil {
			return err
		}
		m.hospitals[hospital.ID] = hospital
	}

	return nil
}

func (m *HospitalManager) getHospitalHandleSymptoms() error {
	for _, hospital := range m.hospitals {
		rows, err := m.db.Query("SELECT symptom_id FROM hospital_handle_symptom WHERE hospital_id = ?", hospital.ID)
		if err != nil {
			return err
		}

		defer rows.Close()

		hospital.Symptoms = make([]*objects.Symptom, 0)

		for rows.Next() {
			var symptomID int
			err := rows.Scan(&symptomID)
			if err != nil {
				return err
			}
			//TODO
		}
	}

	return nil
}

func (m *HospitalManager) getHospitalReviewStat() error {
	for _, hospital := range m.hospitals {
		row := m.db.QueryRow("SELECT average_rating, total_rating, review_count, rating_stability FROM hospital_review_stat WHERE id = ?", hospital.ID)

		stat := objects.ReviewStat{}
		err := row.Scan(&stat.AverageRating, &stat.TotalRating, &stat.ReviewCount, &stat.RatingStability)
		if err != nil {
			return err
		}

		hospital.ReviewStat = stat
	}

	return nil
}

func (m *HospitalManager) getHospitalFacility() error {
	for _, hospital := range m.hospitals {
		row := m.db.QueryRow("SELECT parking_lot FROM hospital_facility WHERE id = ?", hospital.ID)

		var parkingLot int
		err := row.Scan(&parkingLot)
		if err != nil {
			return err
		}

		hospital.Facility = objects.HospitalFacility{HaveParkingLot: parkingLot == 1}
	}

	return nil
}

func (m *HospitalManager) ResetHospitalManager() error {
	m.hospitals = make(map[int]*objects.Hospital)

	err := m.getHospitalBasicInfo()
	if err != nil {
		return err
	}

	err = m.getHospitalHandleSymptoms()
	if err != nil {
		return err
	}

	err = m.getHospitalReviewStat()
	if err != nil {
		return err
	}

	err = m.getHospitalFacility()
	if err != nil {
		return err
	}

	return nil
}

func (m *HospitalManager) GetHospital(id int) *objects.Hospital {
	return m.hospitals[id]
}

func (m *HospitalManager) GetHospitals() ([]*objects.Hospital, error) {
	if len(m.hospitals) == 0 {
		return nil, errors.New("no hospitals")
	}

	hospitals := make([]*objects.Hospital, 0, len(m.hospitals))
	for _, hospital := range m.hospitals {
		hospitals = append(hospitals, hospital)
	}

	return hospitals, nil
}
