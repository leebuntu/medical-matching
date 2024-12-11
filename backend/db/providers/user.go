package providers

import (
	"database/sql"
	"fmt"
	"medical-matching/constants"
	"medical-matching/constants/dto"
	"medical-matching/db"
	"medical-matching/utils"
	"strconv"
	"sync"

	"github.com/google/uuid"
)

type UserProvider struct {
	db *sql.DB
}

var userOnce sync.Once
var userInstance *UserProvider

func GetUserProvider() *UserProvider {
	userOnce.Do(func() {
		db, err := db.GetDBManager().GetDB(constants.UserDB)
		if err != nil {
			return
		}
		userInstance = &UserProvider{db: db}
	})
	return userInstance
}

func (p *UserProvider) getEmailByID(id int, profile *dto.UserProfile) error {
	err := p.db.QueryRow("SELECT email FROM user WHERE id = ?", id).Scan(&profile.Email)
	if err != nil {
		return err
	}
	return nil
}

func (p *UserProvider) getBasicInfoByID(id int, profile *dto.UserProfile) error {
	var profile_url sql.NullString
	var card_id sql.NullString

	err := p.db.QueryRow("SELECT name, profile_image_url, phone_number, home_address, candy, card_id FROM user_profile WHERE id = ?", id).Scan(&profile.Username, &profile_url, &profile.PhoneNumber, &profile.HomeAddress, &profile.Candy, &card_id)
	if err != nil {
		return err
	}

	if profile_url.Valid {
		profile.ProfileURL = profile_url.String
	} else {
		profile.ProfileURL = ""
	}

	if card_id.Valid {
		profile.CardID, err = strconv.Atoi(card_id.String)
		if err != nil {
			return err
		}
	} else {
		profile.CardID = 0
	}

	return nil
}

func (p *UserProvider) GetPriorityByID(id int) ([]int, error) {
	profile := dto.UserProfile{}
	err := p.getPriorArrByID(id, &profile)
	if err != nil {
		return nil, err
	}

	return profile.PriorityOption, nil
}

func (p *UserProvider) getPriorArrByID(id int, profile *dto.UserProfile) error {
	priorityMap := make(map[int]int)

	rows, err := p.db.Query("SELECT priority_id, rank FROM priority WHERE user_id = ?", id)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		fmt.Println(err)
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var priority_id int
		var rank int
		err := rows.Scan(&priority_id, &rank)
		if err != nil {
			return err
		}
		priorityMap[priority_id] = rank
	}

	profile.PriorityOption = utils.SortMapByValueAndGetKeys(priorityMap)

	return nil
}

func (p *UserProvider) GetUserProfile(id int) (*dto.UserProfile, error) {
	profile := dto.UserProfile{}

	err := p.getEmailByID(id, &profile)
	if err != nil {
		return &profile, err
	}

	err = p.getBasicInfoByID(id, &profile)
	if err != nil {
		return &profile, err
	}

	err = p.getPriorArrByID(id, &profile)
	if err != nil {
		return &profile, err
	}

	return &profile, nil
}

func (p *UserProvider) UpdateUserProfile(id int, up *dto.UserProfileUpdate) error {
	_, err := p.db.Exec("DELETE FROM priority WHERE user_id = ?", id)
	if err != nil {
		return err
	}

	for i, priority := range up.PriorityOption {
		_, err := p.db.Exec("INSERT INTO priority (user_id, priority_id, rank) VALUES (?, ?, ?)", id, priority, i+1)
		if err != nil {
			return err
		}
	}

	_, err = p.db.Exec("UPDATE user_profile SET profile_image_url = ?, phone_number = ?, home_address = ?, card_id = ? WHERE id = ?", up.ProfileURL, up.PhoneNumber, up.HomeAddress, up.CardID, id)

	if err != nil {
		return err
	}

	return nil
}

func (p *UserProvider) AddPaymentMethod(userID int, card *dto.PaymentMethod) error {
	cardID := uuid.New().String()

	_, err := p.db.Exec("INSERT INTO payment_method (id, user_id, card_holder_name, card_number, exp_date, cvv) VALUES (?, ?, ?, ?, ?, ?)", cardID, userID, card.CardHolderName, card.CardNumber, card.ExpDate, card.Cvv)
	if err != nil {
		return err
	}

	return nil
}

func (p *UserProvider) GetPaymentMethodList(userID int) ([]*dto.RetrievePaymentMethod, error) {
	rows, err := p.db.Query("SELECT id, card_number, exp_date FROM payment_method WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	paymentMethods := []*dto.RetrievePaymentMethod{}
	for rows.Next() {
		var paymentMethod dto.RetrievePaymentMethod
		err := rows.Scan(&paymentMethod.CardID, &paymentMethod.CardNumber, &paymentMethod.ExpDate)
		if err != nil {
			return nil, err
		}
		paymentMethods = append(paymentMethods, &paymentMethod)
	}

	return paymentMethods, nil
}

func (p *UserProvider) DeletePaymentMethod(userID int, cardID string) error {
	return nil
}
