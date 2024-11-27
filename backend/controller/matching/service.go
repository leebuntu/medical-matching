package matching

import (
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/db"
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

func (m *Matching) getPriority() error {
	db, err := db.GetDBManager().GetDB(constants.UserDB)
	if err != nil {
		return err
	}

	priority, err := user.NewUserService(db).GetPriorityByID(m.userID)
	if err != nil {
		return err
	}

	m.composer = NewComposer(priority)

	return nil
}

func (m *Matching) GetMatchingID() string {
	return m.matchingID
}

func (m *Matching) StartMatching() { // implement only find hospital
	m.state = constants.StartMatching

	if err := m.getPriority(); err != nil {
		m.state = constants.MatchingFailed
		return
	}

}
