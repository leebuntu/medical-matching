package db

import (
	"MedicalMatching/constants"
	"database/sql"
	"sync"

	"github.com/mattn/go-sqlite3"
)

var instance *DBManager
var once sync.Once

type DBManager struct {
	databases map[string]*sql.DB
}

func (m *DBManager) InitDB() {
	for key, value := range constants.DatabaseNames {
		m.addDB(key, value)
	}
}

func GetDBManager() *DBManager {
	once.Do(func() {
		instance = &DBManager{
			databases: make(map[string]*sql.DB),
		}
	})

	return instance
}

func (manager *DBManager) addDB(name, dbPath string) {
	db, err := sql.Open("sqlite3", constants.DBPath+dbPath)
	if err != nil {
		panic(err)
	}

	manager.databases[name] = db
}

func (manager *DBManager) GetDB(name string) (*sql.DB, error) {
	db, ok := manager.databases[name]
	if !ok {
		return nil, sqlite3.ErrError
	}

	return db, nil
}

func (m *DBManager) CloseAll() {
	for _, db := range m.databases {
		db.Close()
	}
}
