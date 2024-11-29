package record

import (
	"database/sql"
	"medical-matching/constants"
	"medical-matching/db"
	"sync"
)

type RecordService struct {
	db *sql.DB
}

var once sync.Once
var instance *RecordService

func GetService() *RecordService {
	once.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.UserDB)
		if err != nil {
			return
		}
		instance = &RecordService{db: db}
	})
	return instance
}
