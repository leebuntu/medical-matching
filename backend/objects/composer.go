package objects

import (
	"math"
	"medical-matching/constants"
	"medical-matching/maps"
	"slices"
	"sort"
)

type Composer struct {
	symptoms       []int
	basisLongitude float64
	basisLatitude  float64
	priority       []int
	methods        []func(hospital *Hospital) (float64, error)
}

func NewComposer(symptoms []int, basisLongitude, basisLatitude float64, priority []int) *Composer {
	instance := &Composer{
		symptoms:       symptoms,
		basisLongitude: basisLongitude,
		basisLatitude:  basisLatitude,
		priority:       priority,
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
	score := float64((100 + (hospital.WaitingPerson * constants.PerWatingPersonScore)))
	if score < 0 {
		return 0, nil
	}
	return score, nil
}

func (c *Composer) calculateDistance(hospital *Hospital) (float64, error) {
	time, err := maps.GetDrivingTimeAsMinutes(c.basisLongitude, c.basisLatitude, hospital.Longitude, hospital.Latitude)
	if err != nil {
		return 0.0, err
	}

	score := 100 + (float64(time) * constants.PerDrivingTimeScore)
	if score < 0 {
		return 0, nil
	}
	return score, nil
}

func (c *Composer) calculateReview(hospital *Hospital) (float64, error) {

	w1 := 4.0
	w2 := 3.0
	w3 := 2.0
	w4 := 1.0

	averageRating := w1 * hospital.ReviewStat.AverageRating
	ratingStability := w2 * (1.0 / (1.0 + hospital.ReviewStat.RatingStability))
	reviewCount := w3 * math.Log10(float64(hospital.ReviewStat.ReviewCount))
	totalRating := 0.0
	if hospital.ReviewStat.ReviewCount > 0 {
		totalRating = w4 * (float64(hospital.ReviewStat.TotalRating) / float64(hospital.ReviewStat.ReviewCount))
	}

	return averageRating + ratingStability + reviewCount + totalRating, nil
}

func (c *Composer) calculateHaveParkingLot(hospital *Hospital) (float64, error) {
	if hospital.Facility.HaveParkingLot == 1 {
		return constants.HaveParkingLotScore, nil
	}
	return 0, nil
}

func (c *Composer) calculateLeastWalk(hospital *Hospital) (float64, error) {
	walkingTime, err := maps.GetPedestrianTimeAsMinutes(c.basisLongitude, c.basisLatitude, hospital.Longitude, hospital.Latitude, "t", "f")
	if err != nil {
		return 0, err
	}
	score := 100 + (walkingTime * constants.PerWalkMinuteScore)
	if score < 0 {
		return 0, nil
	}
	return score, nil
}

func (c *Composer) calculateWeightedScore(scores []float64) []float64 {
	weights := []float64{constants.Weights[1], constants.Weights[2], constants.Weights[3]}
	totalScores := make([]float64, constants.TotalPriority)

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
		totalScores[i] = score * weight
	}

	return totalScores
}

func (c *Composer) getContentRank(scores []float64) []int {
	n := 3

	type ValueIndex struct {
		Value float64
		Index int
	}

	valueIndexes := make([]ValueIndex, len(scores))
	for i, v := range scores {
		valueIndexes[i] = ValueIndex{Value: v, Index: i}
	}

	sort.Slice(valueIndexes, func(i, j int) bool {
		return valueIndexes[i].Value > valueIndexes[j].Value
	})

	topIndexes := make([]int, n)
	for i := 0; i < n && i < len(valueIndexes); i++ {
		if valueIndexes[i].Value > 0 {
			topIndexes[i] = valueIndexes[i].Index + 1
		}
	}

	return topIndexes
}

func (c *Composer) getHospitalScore(hospital *Hospital) (*WeightedScore, error) {
	totalScore := 0.0

	totalScores := make([]float64, constants.TotalPriority)

	for i, method := range c.methods {
		score, err := method(hospital)
		if err != nil {
			return nil, err
		}
		totalScores[i] = score
	}

	totalScores = c.calculateWeightedScore(totalScores)

	for _, score := range totalScores {
		totalScore += score
	}

	return &WeightedScore{
		TotalScore:  totalScore,
		ContentRank: c.getContentRank(totalScores),
	}, nil
}

func (c *Composer) GetHospitalScore(hospital *Hospital) (*WeightedScore, error) {
	if !c.intersectSymptomsWithHospital(hospital) {
		return nil, nil
	}

	score, err := c.getHospitalScore(hospital)
	if err != nil {
		return nil, err
	}

	return score, nil
}
