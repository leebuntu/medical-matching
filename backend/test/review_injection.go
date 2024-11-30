package test

import (
	"database/sql"
	"medical-matching/constants"
	"medical-matching/db"
	"sync"
)

type ReviewInjection struct {
	db *sql.DB
}

var reviewOnce sync.Once
var reviewInstance *ReviewInjection

func GetReviewInjection() *ReviewInjection {
	reviewOnce.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.ReviewDB)
		if err != nil {
			return
		}
		reviewInstance = &ReviewInjection{db: db}
	})

	return reviewInstance
}

func (r *ReviewInjection) alreadyInjected() (bool, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM review").Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *ReviewInjection) InjectReview() error {
	return nil
}
