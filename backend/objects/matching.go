package objects

import (
	"medical-matching/constants"
	"medical-matching/constants/dto"

	"github.com/google/uuid"
)

type Matching struct {
	userID     int
	matchingID string
	symptoms   []int
	composer   *Composer
	state      int
	result     *dto.PoolingResponseCompleted
}

func NewMatching(userID int, symptoms []int) *Matching {
	return &Matching{
		userID:     userID,
		matchingID: uuid.New().String(),
		symptoms:   symptoms,
		state:      constants.BeforeMatching,
	}
}

func (m *Matching) GetState() int {
	return m.state
}

func (m *Matching) GetUserID() int {
	return m.userID
}

func (m *Matching) GetMatchingID() string {
	return m.matchingID
}

func (m *Matching) setComposer(priority []int) {
	m.composer = NewComposer(m.symptoms, priority)
}

func (m *Matching) StartMatching(priority []int, hospitals []*Hospital) {
	m.setComposer(priority)

	m.state = constants.StartMatching

	if m.composer == nil {
		m.state = constants.MatchingFailed
		return
	}

	best := FilteringHospital(hospitals, m.composer)

	m.state = constants.MatchingCompleted

	m.result = &dto.PoolingResponseCompleted{
		State:         m.state,
		HospitalID:    best.HospitalID,
		ContentOption: best.Content,
		WaitingPerson: best.WaitingPerson,
	}
}

func (m *Matching) GetCompleteResult() *dto.PoolingResponseCompleted {
	return m.result
}
