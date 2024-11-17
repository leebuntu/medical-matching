package auth

import (
	"MedicalMatching/constants/dto/auth"
	"database/sql"
	"fmt"
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Login(r *auth.LoginRequest) (int, error) {
	stmt, err := s.db.Prepare("SELECT 1 FROM user WHERE email_address = ? AND hashed_password = ?")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	var userID int
	err = stmt.QueryRow(r.Email, r.HashedPassword).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (s *AuthService) isDuplicateUser(tx *sql.Tx, email string) (bool, error) {
	var exists int
	err := tx.QueryRow("SELECT 1 FROM user WHERE email_address = ?", email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *AuthService) Register(r *auth.RegisterRequest) (bool, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return false, err
	}

	defer tx.Rollback()

	if exists, err := s.isDuplicateUser(tx, r.Email); exists {
		return false, err
	}

	result, err := tx.Exec("INSERT INTO user (email_address, hashed_password) VALUES (?, ?)", r.Email, r.HashedPassword)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return false, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return false, err
	}

	_, err = tx.Exec("INSERT INTO user_profile (id, name, phone_number, home_address, postal_code) VALUES (?, ?, ?, ?, ?)", userID, r.Username, r.PhoneNumber, r.HomeAddress, r.PostalCode)
	if err != nil {
		fmt.Println("Error inserting user profile:", err)
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}
