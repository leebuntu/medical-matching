package matching

import (
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/db/hospital"
	"medical-matching/db/user"

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

func (m *Matching) setComposer() error {
	priority, err := user.GetService().GetPriorityByID(m.userID)
	if err != nil {
		return err
	}

	m.composer = NewComposer(priority)

	return nil
}

func (m *Matching) GetMatchingID() string {
	return m.matchingID
}

func (m *Matching) StartMatching() {
	m.state = constants.StartMatching

	if err := m.setComposer(); err != nil {
		m.state = constants.MatchingFailed
		return
	}

	// TODO
	mm := hospital.GetHospitalManager()
	hospitals, err := mm.GetHospitals()
	if err != nil {
		m.state = constants.MatchingFailed
		return
	}

	best := FilteringHospital(hospitals, m.composer)

	m.result = &dto.PoolingResponseCompleted{
		HospitalID: best.HospitalID,
	}
}

func (m *Matching) GetCompleteResult() *dto.PoolingResponseCompleted {
	return m.result
}
