package matching

import (
	"errors"
	"medical-matching/objects"
	"sync"
)

type MatchingManager struct {
	matchings map[string]*objects.Matching
}

var once sync.Once
var instance *MatchingManager

func GetMatchingManager() *MatchingManager {
	once.Do(func() {
		instance = &MatchingManager{
			matchings: make(map[string]*objects.Matching),
		}
	})

	return instance
}

func (m *MatchingManager) GetMatching(matchingID string) (*objects.Matching, error) {
	matching, ok := m.matchings[matchingID]
	if !ok {
		return nil, errors.New("matching not found")
	}
	return matching, nil
}

func (m *MatchingManager) CreateMatching(userID int, symptoms []int) (*objects.Matching, error) {
	// TODO: check limit

	matching := objects.NewMatching(userID, symptoms)

	m.matchings[matching.GetMatchingID()] = matching

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
