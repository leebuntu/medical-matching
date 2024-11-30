package objects

import (
	"medical-matching/constants"
	"medical-matching/constants/dto"

	"github.com/google/uuid"
)

type Matching struct {
	userID     int
	matchingID string
	context    *dto.MatchingRequest
	composer   *Composer
	state      int
	result     *dto.PoolingResponseCompleted
}

func NewMatching(userID int, context *dto.MatchingRequest) *Matching {
	return &Matching{
		userID:     userID,
		matchingID: uuid.New().String(),
		context:    context,
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

func (m *Matching) SetComposer(priority []int) {
	m.composer = NewComposer(priority)
}

func (m *Matching) StartMatching(hospitals []*Hospital) {
	m.state = constants.StartMatching

	if m.composer == nil {
		m.state = constants.MatchingFailed
		return
	}

	best := FilteringHospital(hospitals, m.composer)

	m.result = &dto.PoolingResponseCompleted{
		HospitalID:    best.HospitalID,
		ContentOption: best.Content,
		WaitingPerson: best.WaitingPerson,
	}

	m.state = constants.MatchingCompleted
}

func (m *Matching) GetCompleteResult() *dto.PoolingResponseCompleted {
	return m.result
}
