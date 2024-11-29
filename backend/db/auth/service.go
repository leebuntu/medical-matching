package auth

import (
	"database/sql"
	"medical-matching/constants"
	"medical-matching/db"
	"sync"
)

type AuthService struct {
	db *sql.DB
}

var once sync.Once
var authInstance *AuthService

func GetService() *AuthService {
	once.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.UserDB)
		if err != nil {
			return
		}
		authInstance = &AuthService{db: db}
	})
	return authInstance
}
