package objects

import (
	"medical-matching/constants"
	"slices"
)

type Composer struct {
	symptoms []int
	priority []int
	methods  []func(hospital *Hospital) (float64, error)
}

func NewComposer(symptoms []int, priority []int) *Composer {
	instance := &Composer{
		symptoms: symptoms,
		priority: priority,
	}
	instance.init()
	return instance
}

func (c *Composer) init() {
	c.methods = []func(hospital *Hospital) (float64, error){
		c.calculateWaiting,
		c.calculateDistance,
		c.calculateReview,
		c.calculateHaveParkingLot,
		c.calculateLeastWalk,
	}

}

func (c *Composer) intersectSymptomsWithHospital(hospital *Hospital) bool {
	exist := false

	for _, symptom := range c.symptoms {
		if slices.Contains(hospital.HandleSymptoms, symptom) {
			exist = true
			break
		}
	}

	return exist
}

func (c *Composer) calculateWaiting(hospital *Hospital) (float64, error) {
	return float64((100 - (hospital.WaitingPerson * constants.PerWatingPersonScore))), nil
}

func (c *Composer) calculateDistance(hospital *Hospital) (float64, error) {
	// TODO: using naver api or other api
	return 0.0, nil
}

func (c *Composer) calculateReview(hospital *Hospital) (float64, error) {
	// TODO: calculate review but using random number maybe?
	return 0.0, nil
}

func (c *Composer) calculateHaveParkingLot(hospital *Hospital) (float64, error) {
	if hospital.Facility.HaveParkingLot == 1 {
		return constants.HaveParkingLotScore, nil
	}
	return 0, nil
}

func (c *Composer) calculateLeastWalk(hospital *Hospital) (float64, error) {
	// TODO: using naver api or other api
	return 0.0, nil
}

func (c *Composer) calculateWeightedScore(scores []float64) float64 {
	weights := []float64{constants.Weights[1], constants.Weights[2], constants.Weights[3]}
	totalScore := 0.0

	priorityWeight := make(map[int]float64)
	for i, priority := range c.priority {
		if i < len(weights) {
			priorityWeight[priority] = weights[i]
		}
	}

	for i, score := range scores {
		weight, exists := priorityWeight[i+1]
		if !exists {
			weight = constants.Weights[4]
		}
		totalScore += score * weight
	}

	return totalScore
}

func (c *Composer) getHospitalScore(hospital *Hospital) (float64, error) {
	totalScore := 0.0

	totalScores := make([]float64, constants.TotalPriority)

	for i, method := range c.methods {
		score, err := method(hospital)
		if err != nil {
			return 0, err
		}
		totalScores[i] = score
	}

	totalScore = c.calculateWeightedScore(totalScores)

	return totalScore, nil
}

func (c *Composer) GetHospitalScore(hospital *Hospital) (float64, error) {
	if !c.intersectSymptomsWithHospital(hospital) {
		return 0, nil
	}

	score, err := c.getHospitalScore(hospital)
	if err != nil {
		return 0, err
	}

	return score, nil
}
