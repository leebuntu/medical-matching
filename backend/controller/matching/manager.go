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

func (m *MatchingManager) RemoveMatching(matchingID string) error {
	_, ok := m.matchings[matchingID]
	if !ok {
		return errors.New("matching not found")
	}
	delete(m.matchings, matchingID)
	return nil
}

func (m *MatchingManager) GetAllMatching(userID int) []string {
	matchingIDs := make([]string, 0)
	for _, matching := range m.matchings {
		if matching.GetUserID() == userID {
			matchingIDs = append(matchingIDs, matching.GetMatchingID())
		}
	}
	return matchingIDs
}
