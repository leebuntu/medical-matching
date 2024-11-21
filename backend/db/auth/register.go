package auth

import (
	"database/sql"
	"medical-matching/constants/dto"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) isDuplicateUser(tx *sql.Tx, email string) (bool, error) {
	var exists int
	err := tx.QueryRow("SELECT 1 FROM user WHERE email = ?", email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *AuthService) Register(r *dto.RegisterRequest) (bool, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return true, err
	}

	defer tx.Rollback()

	if exists, err := s.isDuplicateUser(tx, r.Email); exists {
		return false, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return true, err
	}
	r.Password = string(hashedPassword)

	result, err := tx.Exec("INSERT INTO user (email, hashed_password) VALUES (?, ?)", r.Email, r.Password)
	if err != nil {
		return true, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return true, err
	}

	_, err = tx.Exec("INSERT INTO user_profile (id, name, phone_number, home_address, postal_code) VALUES (?, ?, ?, ?, ?)", userID, r.Username, r.PhoneNumber, r.HomeAddress, r.PostalCode)
	if err != nil {
		return true, err
	}

	err = tx.Commit()
	if err != nil {
		return true, err
	}

	return true, nil
}
