package matching

import (
	"errors"
	"medical-matching/constants/dto"
	"sync"
)

type MatchingManager struct {
	matchings map[int]*Matching
}

var once sync.Once
var instance *MatchingManager

func GetMatchingManager() *MatchingManager {
	once.Do(func() {
		instance = &MatchingManager{
			matchings: make(map[int]*Matching),
		}
	})

	return instance
}

func (m *MatchingManager) alreadyMatched(userID int) bool {
	_, ok := m.matchings[userID]
	return ok
}

func (m *MatchingManager) CreateMatching(userID int, context dto.MatchingRequest) error {
	matching := &Matching{
		userID:  userID,
		context: &context,
	}

	if m.alreadyMatched(userID) {
		return errors.New("already matched")
	}

	m.matchings[userID] = matching

	return nil
}

func (m *MatchingManager) StartMatching(userID int) error {
	return m.matchings[userID].StartMatching()
}
