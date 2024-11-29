package matching

import (
	"errors"
	"medical-matching/constants/dto"
	"sync"
)

type MatchingManager struct {
	matchings map[string]*Matching
}

var once sync.Once
var instance *MatchingManager

func GetMatchingManager() *MatchingManager {
	once.Do(func() {
		instance = &MatchingManager{
			matchings: make(map[string]*Matching),
		}
	})

	return instance
}

func (m *MatchingManager) GetMatching(matchingID string) (*Matching, error) {
	matching, ok := m.matchings[matchingID]
	if !ok {
		return nil, errors.New("matching not found")
	}
	return matching, nil
}

func (m *MatchingManager) CreateMatching(userID int, context *dto.MatchingRequest) (*Matching, error) {
	// TODO: check limit

	matching := NewMatching(userID, context)

	m.matchings[matching.matchingID] = matching

	return matching, nil
}

func (m *MatchingManager) RemoveMatching(matchingID string) {
	delete(m.matchings, matchingID)
}
