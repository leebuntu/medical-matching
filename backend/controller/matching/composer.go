package matching

import (
	"medical-matching/constants/dto"
	"medical-matching/constants/objects"
)

type Composer struct {
	priority []int
	methods  []func(hospital *objects.Hospital, weight float64) (float64, error)
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
	c.methods = []func(hospital *objects.Hospital, weight float64) (float64, error){
		nil,
		c.calculateWaiting,
		c.calculateDistance,
		c.calculateReview,
		c.calculateHaveParkingLot,
		c.calculateLeastWalk,
	}

	c.weights = make(map[int]float64)
	c.weights = dto.Weights
	for i := len(c.weights) + 1; i <= dto.TotalPriority; i++ {
		c.weights[i] = 1.0
	}

	c.orders = make([]int, dto.TotalPriority)
	copy(c.orders, c.priority)

	exists := make(map[int]bool)
	for _, val := range c.priority {
		exists[val] = true
	}

	index := len(c.priority)
	for i := 1; i <= dto.TotalPriority; i++ {
		if !exists[i] {
			c.orders[index] = i
			index++
		}
	}
}

func (c *Composer) calculateWaiting(hospital *objects.Hospital, weight float64) (float64, error) {
	return float64((100 - (hospital.WaitingPerson * dto.PerWatingPersonScore))) * weight, nil
}

func (c *Composer) calculateDistance(hospital *objects.Hospital, weight float64) (float64, error) {
	// TODO: using naver api or other api
	return 0.0, nil
}

func (c *Composer) calculateReview(hospital *objects.Hospital, weight float64) (float64, error) {
	// TODO: calculate review but using random number maybe?
	return 0.0, nil
}

func (c *Composer) calculateHaveParkingLot(hospital *objects.Hospital, weight float64) (float64, error) {
	if hospital.Facility.HaveParkingLot {
		return dto.HaveParkingLotScore * weight, nil
	}
	return 0, nil
}

func (c *Composer) calculateLeastWalk(hospital *objects.Hospital, weight float64) (float64, error) {
	// TODO: using naver api or other api
	return 0.0, nil
}

func (c *Composer) getHospitalScore(hospital *objects.Hospital) (float64, error) {
	totalScore := 0.0

	for index, order := range c.orders {
		score, err := c.methods[order](hospital, c.weights[index])
		if err != nil {
			return 0, err
		}
		totalScore += score
	}

	return totalScore, nil
}

func (c *Composer) GetHospitalScore(hospital *objects.Hospital) (float64, error) {
	score, err := c.getHospitalScore(hospital)
	if err != nil {
		return 0, err
	}

	return score, nil
}
