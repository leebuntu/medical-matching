package user

import (
	"MedicalMatching/constants/dto"
	"database/sql"
)

type UserService struct {
	db      *sql.DB
	profile dto.UserProfile
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db, profile: dto.UserProfile{}}
}
