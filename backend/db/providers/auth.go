package providers

import (
	"database/sql"
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/db"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type AuthProvider struct {
	db *sql.DB
}

var authOnce sync.Once
var authInstance *AuthProvider

func GetAuthProvider() *AuthProvider {
	authOnce.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.UserDB)
		if err != nil {
			return
		}
		authInstance = &AuthProvider{db: db}
	})
	return authInstance
}

func (p *AuthProvider) Login(r *dto.LoginRequest) (int, error) {
	var userID int
	var hashedPassword string

	err := p.db.QueryRow("SELECT id, hashed_password FROM user WHERE email = ?", r.Email).Scan(&userID, &hashedPassword)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(r.Password))
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (p *AuthProvider) isDuplicateUser(tx *sql.Tx, email string) (bool, error) {
	var exists int
	err := tx.QueryRow("SELECT 1 FROM user WHERE email = ?", email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *AuthProvider) Register(r *dto.RegisterRequest) (bool, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return true, err
	}

	defer tx.Rollback()

	if exists, err := p.isDuplicateUser(tx, r.Email); exists {
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
