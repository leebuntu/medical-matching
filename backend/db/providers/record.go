package providers

import (
	"database/sql"
	"medical-matching/constants"
	"medical-matching/db"
	"sync"
)

type RecordProvider struct {
	db *sql.DB
}

var recordOnce sync.Once
var recordInstance *RecordProvider

func GetRecordProvider() *RecordProvider {
	recordOnce.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.UserDB)
		if err != nil {
			return
		}
		recordInstance = &RecordProvider{db: db}
	})
	return recordInstance
}
