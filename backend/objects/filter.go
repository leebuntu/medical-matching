package objects

import (
	"fmt"
	"sync"
)

type ScoredHospital struct {
	HospitalID    int
	Score         float64
	Content       []string
	WaitingPerson int
}

func FilteringHospital(hospitals []*Hospital, composer *Composer) *ScoredHospital {
	resultChan := make(chan ScoredHospital, len(hospitals))

	var wait sync.WaitGroup

	for _, hospital := range hospitals {
		wait.Add(1)
		go func(hospital *Hospital) {
			defer wait.Done()
			score, err := composer.GetHospitalScore(hospital)
			if err != nil {
				fmt.Println(err)
				resultChan <- ScoredHospital{HospitalID: hospital.ID, Score: 0, WaitingPerson: hospital.WaitingPerson}
				return
			}
			resultChan <- ScoredHospital{HospitalID: hospital.ID, Score: score, WaitingPerson: hospital.WaitingPerson}
		}(hospital)
	}

	wait.Wait()
	close(resultChan)

	best := &ScoredHospital{Score: 0}

	for scores := range resultChan {
		if best == nil || scores.Score > best.Score {
			best = &scores
		}
	}

	return best
}
