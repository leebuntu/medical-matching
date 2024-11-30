package providers

import (
	"database/sql"
	"medical-matching/constants"
	"medical-matching/db"
	"medical-matching/objects"
	"sync"
)

type HospitalProvider struct {
	db *sql.DB
}

var hospitalOnce sync.Once
var hospitalInstance *HospitalProvider

func GetHospitalProvider() *HospitalProvider {
	hospitalOnce.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.HospitalDB)
		if err != nil {
			return
		}
		hospitalInstance = &HospitalProvider{db: db}
	})
	return hospitalInstance
}

func (m *HospitalProvider) getHospitalBasicInfo(hospitals *[]*objects.Hospital) error {
	rows, err := m.db.Query("SELECT id, name, owner_name, address, postal_code, longitude, latitude, contact_phone_number FROM hospital")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		hospital := &objects.Hospital{}
		err := rows.Scan(&hospital.ID, &hospital.Name, &hospital.OwnerName, &hospital.Address, &hospital.PostalCode, &hospital.Longitude, &hospital.Latitude, &hospital.ContactPhoneNumber)
		if err != nil {
			return err
		}
		*hospitals = append(*hospitals, hospital)
	}

	return nil
}

func (m *HospitalProvider) getHospitalHandleSymptoms(hospitals *[]*objects.Hospital) error {
	for _, hospital := range *hospitals {
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
			hospital.Symptoms = append(hospital.Symptoms, &objects.Symptom{ID: symptomID})
		}
	}

	return nil
}

func (m *HospitalProvider) getHospitalReviewStat(hospitals *[]*objects.Hospital) error {
	for _, hospital := range *hospitals {
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

func (p *HospitalProvider) getHospitalFacility(hospitals *[]*objects.Hospital) error {
	for _, hospital := range *hospitals {
		row := p.db.QueryRow("SELECT parking_lot FROM hospital_facility WHERE id = ?", hospital.ID)

		var parkingLot int
		err := row.Scan(&parkingLot)
		if err != nil {
			return err
		}

		hospital.Facility = objects.HospitalFacility{HaveParkingLot: parkingLot}
	}

	return nil
}

func (p *HospitalProvider) FetchHospitals() ([]*objects.Hospital, error) {
	hospitals := make([]*objects.Hospital, 0)

	err := p.getHospitalBasicInfo(&hospitals)
	if err != nil {
		return nil, err
	}

	err = p.getHospitalHandleSymptoms(&hospitals)
	if err != nil {
		return nil, err
	}

	err = p.getHospitalReviewStat(&hospitals)
	if err != nil {
		return nil, err
	}

	err = p.getHospitalFacility(&hospitals)
	if err != nil {
		return nil, err
	}

	return hospitals, nil
}
