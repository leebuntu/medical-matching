package db

import "database/sql"

type Service interface {
	NewService(db *sql.DB) Service
}
