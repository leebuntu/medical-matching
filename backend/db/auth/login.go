package auth

import (
	"medical-matching/constants/dto"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(r *dto.LoginRequest) (int, error) {
	var userID int
	var hashedPassword string

	err := s.db.QueryRow("SELECT id, hashed_password FROM user WHERE email = ?", r.Email).Scan(&userID, &hashedPassword)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(r.Password))
	if err != nil {
		return 0, err
	}

	return userID, nil
}
