package user

import (
	"database/sql"
	"medical-matching/constants"
	"medical-matching/db"
	"sync"
)

type UserService struct {
	db *sql.DB
}

var once sync.Once
var userInstance *UserService

func GetService() *UserService {
	once.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.UserDB)
		if err != nil {
			return
		}
		userInstance = &UserService{db: db}
	})
	return userInstance
}
