package matching

import (
	"medical-matching/constants/objects"
	"sync"
)

func FilteringHospital(hospitals []*objects.Hospital, composer *Composer) *objects.ScoredHospital {
	// TODO: Filtering with go routine and channel
	resultChan := make(chan objects.ScoredHospital, len(hospitals))

	var wait sync.WaitGroup

	for _, hospital := range hospitals {
		wait.Add(1)
		go func(hospital *objects.Hospital) {
			defer wait.Done()
			score, err := composer.GetHospitalScore(hospital)
			if err != nil {
				resultChan <- objects.ScoredHospital{HospitalID: hospital.ID, Score: 0}
				return
			}
			resultChan <- objects.ScoredHospital{HospitalID: hospital.ID, Score: score}
		}(hospital)
	}

	wait.Wait()
	close(resultChan)

	best := &objects.ScoredHospital{Score: 0}

	for scores := range resultChan {
		if best == nil || scores.Score > best.Score {
			best = &scores
		}
	}

	return best
}
