package sets

import "database/sql"

type HospitalInjection struct {
	db *sql.DB
}

func NewHospitalInjection(db *sql.DB) *HospitalInjection {
	return &HospitalInjection{db: db}
}

func (h *HospitalInjection) alreadyInjected() (bool, error) {
	var count int
	err := h.db.QueryRow("SELECT COUNT(*) FROM hospital").Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (h *HospitalInjection) InjectTestHospital() error {
	alreadyInjected, err := h.alreadyInjected()
	if err != nil {
		return err
	}
	if alreadyInjected {
		return nil
	}

	_, err = h.db.Exec("INSERT INTO hospital (name, owner_name, address, postal_code, contact_phone_number) VALUES (?, ?, ?, ?, ?)", "test", "test", "test", "test", "test")

	if err != nil {
		return err
	}

	return nil
}
