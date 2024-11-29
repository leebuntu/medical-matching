package sets

import (
	"database/sql"
	"medical-matching/constants"
	"medical-matching/db"
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

func (h *HospitalInjection) InjectHospital() error {
	alreadyInjected, err := h.alreadyInjected()
	if err != nil {
		return err
	}
	if alreadyInjected {
		return nil
	}

	return nil
}
