package matching

import (
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/db"
	"medical-matching/db/user"
)

type Matching struct {
	userID   int
	context  *dto.MatchingRequest
	composer *Composer
	state    int
}

func NewMatching(userID int, context *dto.MatchingRequest) *Matching {
	return &Matching{
		userID:  userID,
		context: context,
		state:   dto.BeforeMatching,
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

func (m *Matching) StartMatching() error { // implement only find hospital
	m.state = dto.StartMatching

	if err := m.getPriority(); err != nil {
		return err
	}

	return nil
}
