package hospital

import (
	"errors"
	"medical-matching/objects"
	"sync"
)

type SymptomManager struct {
	symptoms map[int]*objects.Symptom
}

var symptomOnce sync.Once
var symptomInstance *SymptomManager

func GetSymptomManager() *SymptomManager {
	symptomOnce.Do(func() {
		symptomInstance = &SymptomManager{}
	})

	return symptomInstance
}

func (m *SymptomManager) ResetSymptomManager(symptoms []*objects.Symptom) error {
	if len(symptoms) == 0 {
		return errors.New("no symptoms")
	}

	m.symptoms = make(map[int]*objects.Symptom)

	for _, symptom := range symptoms {
		m.symptoms[symptom.ID] = symptom
	}

	return nil
}

func (m *SymptomManager) GetSymptom(id int) *objects.Symptom {
	return m.symptoms[id]
}
