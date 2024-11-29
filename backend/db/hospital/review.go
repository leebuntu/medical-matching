package hospital

import (
	"database/sql"
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/db"
	"medical-matching/db/user"
	"sync"
)

type ReviewManager struct {
	db *sql.DB
}

var reviewOnce sync.Once
var reviewInstance *ReviewManager

func GetReviewManager() *ReviewManager {
	reviewOnce.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.ReviewDB)
		if err != nil {
			return
		}
		reviewInstance = &ReviewManager{db: db}
	})
	return reviewInstance
}

func (m *ReviewManager) GetReview(hospitalID string, page int) ([]*dto.ReviewStat, error) {
	rows, err := m.db.Query("SELECT id, user_id, timestamp, score, context FROM review WHERE hospital_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?", hospitalID, constants.ReviewPerPage, (page-1)*constants.ReviewPerPage)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	reviews := make([]*dto.ReviewStat, 0)

	for rows.Next() {
		review := &dto.ReviewStat{}
		var reviewID int
		var userID int
		err := rows.Scan(&reviewID, &userID, &review.VisitedDate, &review.Rating, &review.ReviewContext)
		if err != nil {
			return nil, err
		}

		usr, err := user.GetService().GetUserProfile(userID)
		if err != nil {
			return nil, err
		}
		review.ProfileURL = usr.ProfileURL
		review.ProfileName = usr.Username

		r, err := m.db.Query("SELECT photo_url FROM photo WHERE review_id = ?", reviewID)
		if err != nil {
			return nil, err
		}

		review.ReviewPhoto = make([]string, 0)
		for r.Next() {
			var photoURL string
			err := r.Scan(&photoURL)
			if err != nil {
				return nil, err
			}
			review.ReviewPhoto = append(review.ReviewPhoto, photoURL)
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}
