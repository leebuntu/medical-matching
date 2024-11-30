package hospital

import (
	"errors"
	"medical-matching/objects"
	"sync"
)

type HospitalManager struct {
	hospitals map[int]*objects.Hospital
}

var hospitalOnce sync.Once
var hospitalInstance *HospitalManager

func GetHospitalManager() *HospitalManager {
	hospitalOnce.Do(func() {
		hospitalInstance = &HospitalManager{}
	})

	return hospitalInstance
}

func (m *HospitalManager) ResetHospitalManager(hospitals []*objects.Hospital) error {
	if len(hospitals) == 0 {
		return errors.New("no hospitals")
	}

	m.hospitals = make(map[int]*objects.Hospital)

	for _, hospital := range hospitals {
		m.hospitals[hospital.ID] = hospital
	}

	return nil
}

func (m *HospitalManager) GetHospital(id int) *objects.Hospital {
	return m.hospitals[id]
}

func (m *HospitalManager) GetHospitals(longitude, latitude float64, radius float64) ([]*objects.Hospital, error) {
	if len(m.hospitals) == 0 {
		return nil, errors.New("no hospitals")
	}

	hospitals := make([]*objects.Hospital, 0, len(m.hospitals))
	for _, hospital := range m.hospitals {
		hospitals = append(hospitals, hospital)
	}
	//TODO

	return hospitals, nil
}
