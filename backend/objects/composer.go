package objects

import (
	"medical-matching/constants"
)

type Composer struct {
	priority []int
	methods  []func(hospital *Hospital, weight float64) (float64, error)
	weights  map[int]float64
	orders   []int
}

func NewComposer(priority []int) *Composer {
	instance := &Composer{
		priority: priority,
	}
	instance.init()
	return instance
}

func (c *Composer) init() {
	c.methods = []func(hospital *Hospital, weight float64) (float64, error){
		nil,
		c.calculateWaiting,
		c.calculateDistance,
		c.calculateReview,
		c.calculateHaveParkingLot,
		c.calculateLeastWalk,
	}

	c.weights = make(map[int]float64)
	c.weights = constants.Weights
	for i := len(c.weights) + 1; i <= constants.TotalPriority; i++ {
		c.weights[i] = 1.0
	}

	c.orders = make([]int, constants.TotalPriority)
	copy(c.orders, c.priority)

	exists := make(map[int]bool)
	for _, val := range c.priority {
		exists[val] = true
	}

	index := len(c.priority)
	for i := 1; i <= constants.TotalPriority; i++ {
		if !exists[i] {
			c.orders[index] = i
			index++
		}
	}
}

func (c *Composer) calculateWaiting(hospital *Hospital, weight float64) (float64, error) {
	return float64((100 - (hospital.WaitingPerson * constants.PerWatingPersonScore))) * weight, nil
}

func (c *Composer) calculateDistance(hospital *Hospital, weight float64) (float64, error) {
	// TODO: using naver api or other api
	return 0.0, nil
}

func (c *Composer) calculateReview(hospital *Hospital, weight float64) (float64, error) {
	// TODO: calculate review but using random number maybe?
	return 0.0, nil
}

func (c *Composer) calculateHaveParkingLot(hospital *Hospital, weight float64) (float64, error) {
	if hospital.Facility.HaveParkingLot == 1 {
		return constants.HaveParkingLotScore * weight, nil
	}
	return 0, nil
}

func (c *Composer) calculateLeastWalk(hospital *Hospital, weight float64) (float64, error) {
	// TODO: using naver api or other api
	return 0.0, nil
}

func (c *Composer) getHospitalScore(hospital *Hospital) (float64, error) {
	totalScore := 0.0

	for i, order := range c.orders {
		score, err := c.methods[order](hospital, c.weights[i])
		if err != nil {
			return 0, err
		}
		totalScore += score
	}

	return totalScore, nil
}

func (c *Composer) GetHospitalScore(hospital *Hospital) (float64, error) {
	score, err := c.getHospitalScore(hospital)
	if err != nil {
		return 0, err
	}

	return score, nil
}
