package providers

import (
	"database/sql"
	"errors"
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/controller/hospital"
	"medical-matching/db"
	"sync"
	"time"
)

type RecordProvider struct {
	db *sql.DB
}

var recordOnce sync.Once
var recordInstance *RecordProvider

func GetRecordProvider() *RecordProvider {
	recordOnce.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.UserDB)
		if err != nil {
			return
		}
		recordInstance = &RecordProvider{db: db}
	})
	return recordInstance
}

func (p *RecordProvider) GetRecordList(userID int, page int) ([]*dto.MedicalRecord, error) {
	offset := (page - 1) * constants.RecordPerPage
	rows, err := p.db.Query("SELECT id, hospital_id, timestamp, hospital_name, doctor_name, notes, symptom FROM medical_record WHERE user_id = ? ORDER BY timestamp DESC LIMIT ? OFFSET ?", userID, constants.RecordPerPage, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records := []*dto.MedicalRecord{}
	for rows.Next() {
		var record dto.MedicalRecord
		var notes sql.NullString
		if err := rows.Scan(&record.ID, &record.HospitalID, &record.Timestamp, &record.HospitalName, &record.DoctorName, &notes, &record.Symptom); err != nil {
			return nil, err
		}
		if notes.Valid {
			record.Notes = notes.String
		}
		records = append(records, &record)
	}

	return records, nil
}

func (p *RecordProvider) UpdateRecordNotes(recordID int, userID int, notes string) error {
	_, err := p.db.Exec("UPDATE medical_record SET notes = ? WHERE id = ? AND user_id = ?", notes, recordID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (p *RecordProvider) AddRecord(userID int, hospitalID int, symptom string) error {
	hm := hospital.GetHospitalManager()
	hospital := hm.GetHospital(hospitalID)
	if hospital == nil {
		return errors.New("hospital not found")
	}

	_, err := p.db.Exec("INSERT INTO medical_record (user_id, hospital_id, timestamp, hospital_name, doctor_name, symptom) VALUES (?, ?, ?, ?, ?, ?)", userID, hospitalID, time.Now(), hospital.Name, hospital.OwnerName, symptom)
	if err != nil {
		return err
	}
	return nil
}
