package dto

import "time"

type UpdateRecordNotesRequest struct {
	Notes string `json:"notes"`
}

type RecordListResponse struct {
	Records []*MedicalRecord `json:"records"`
}

type MedicalRecord struct {
	ID           int       `json:"id"`
	Timestamp    time.Time `json:"timestamp"`
	HospitalID   int       `json:"hospital_id"`
	HospitalName string    `json:"hospital_name"`
	DoctorName   string    `json:"doctor_name"`
	Notes        string    `json:"notes"`
	Symptom      string    `json:"symptom"`
}
