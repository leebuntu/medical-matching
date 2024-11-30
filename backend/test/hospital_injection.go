package test

import (
	"database/sql"
	"encoding/csv"
	"medical-matching/constants"
	"medical-matching/db"
	"os"
	"strconv"
	"strings"
	"sync"
)

type HospitalInjection struct {
	db *sql.DB
}

var hospitalOnce sync.Once
var hospitalInstance *HospitalInjection

func GetHospitalInjection() *HospitalInjection {
	hospitalOnce.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.HospitalDB)
		if err != nil {
			return
		}
		hospitalInstance = &HospitalInjection{db: db}
	})
	return hospitalInstance
}

func (h *HospitalInjection) alreadyInjected() (bool, error) {
	var count int
	err := h.db.QueryRow("SELECT COUNT(*) FROM hospital").Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (h *HospitalInjection) injectHospitalFacility(id int, parkingLot int) error {
	_, err := h.db.Exec("INSERT INTO hospital_facility (id, parking_lot) VALUES (?, ?)", id, parkingLot)
	if err != nil {
		return err
	}

	return nil
}

func (h *HospitalInjection) injectHospitalOpenTime() error {
	// TODO: Implement
	return nil
}

func (h *HospitalInjection) injectHospitalReviewStat(id int, averageRating float64, totalRating int, reviewCount int, ratingStability float64) error {
	_, err := h.db.Exec("INSERT INTO hospital_review_stat (id, average_rating, total_rating, review_count, rating_stability) VALUES (?, ?, ?, ?, ?)", id, averageRating, totalRating, reviewCount, ratingStability)
	if err != nil {
		return err
	}

	return nil
}

func (h *HospitalInjection) injectHospitalHandleSymptom(id int, symptoms []int) error {
	for _, symptom := range symptoms {
		_, err := h.db.Exec("INSERT INTO hospital_handle_symptom (hospital_id, symptom_id) VALUES (?, ?)", id, symptom)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *HospitalInjection) injectHospitalBasic(name string, ownerName string, address string, postalCode string, phoneNumber string) (int, error) {
	result, err := h.db.Exec("INSERT INTO hospital (name, owner_name, address, postal_code, contact_phone_number) VALUES (?, ?, ?, ?, ?)", name, ownerName, address, postalCode, phoneNumber)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	return int(lastID), err
}

func (h *HospitalInjection) InjectHospital() error {
	alreadyInjected, err := h.alreadyInjected()
	if err != nil {
		return err
	}
	if alreadyInjected {
		return nil
	}

	file, err := os.Open(constants.DBPath + constants.TestData)
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

		id, err := h.injectHospitalBasic(record[0], record[1], record[2], record[3], record[4])
		if err != nil {
			return err
		}

		avg, _ := strconv.ParseFloat(record[5], 64)
		tot, _ := strconv.Atoi(record[6])
		cnt, _ := strconv.Atoi(record[7])
		stb, _ := strconv.ParseFloat(record[8], 64)

		err = h.injectHospitalReviewStat(id, avg, tot, cnt, stb)
		if err != nil {
			return err
		}

		park, _ := strconv.Atoi(record[9])
		err = h.injectHospitalFacility(id, park)
		if err != nil {
			return err
		}

		symptoms := strings.Split(record[10], ",")
		symptomsInt := make([]int, len(symptoms))
		for i, s := range symptoms {
			symptomsInt[i], _ = strconv.Atoi(s)
		}
		err = h.injectHospitalHandleSymptom(id, symptomsInt)
		if err != nil {
			return err
		}
	}

	return nil
}
