package user

import (
	"database/sql"
	"fmt"
	"medical-matching/constants/dto"
	"medical-matching/utils"
)

func (s *UserService) getEmailByID(id int, profile *dto.UserProfile) error {
	err := s.db.QueryRow("SELECT email FROM user WHERE id = ?", id).Scan(&profile.Email)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) getBasicInfoByID(id int, profile *dto.UserProfile) error {
	var profile_url sql.NullString
	var payment_id sql.NullString

	err := s.db.QueryRow("SELECT name, profile_image_url, phone_number, home_address, candy, card_id FROM user_profile WHERE id = ?", id).Scan(&profile.Username, &profile_url, &profile.PhoneNumber, &profile.HomeAddress, &profile.Candy, &payment_id)
	if err != nil {
		return err
	}

	if profile_url.Valid {
		profile.ProfileURL = profile_url.String
	} else {
		profile.ProfileURL = ""
	}

	if payment_id.Valid {
		profile.CardID = payment_id.String
	} else {
		profile.CardID = ""
	}

	return nil
}

func (s *UserService) GetPriorityByID(id int) ([]int, error) {
	profile := dto.UserProfile{}
	err := s.getPriorArrByID(id, &profile)
	if err != nil {
		return nil, err
	}

	return profile.PriorityOption, nil
}

func (s *UserService) getPriorArrByID(id int, profile *dto.UserProfile) error {
	priorityMap := make(map[int]int)

	rows, err := s.db.Query("SELECT priority_id, rank FROM priority WHERE user_id = ?", id)
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

func (s *UserService) GetUserProfile(id int) (*dto.UserProfile, error) {
	profile := dto.UserProfile{}

	err := s.getEmailByID(id, &profile)
	if err != nil {
		return &profile, err
	}

	err = s.getBasicInfoByID(id, &profile)
	if err != nil {
		return &profile, err
	}

	err = s.getPriorArrByID(id, &profile)
	if err != nil {
		return &profile, err
	}

	return &profile, nil
}

func (s *UserService) UpdateUserProfile(id int, up *dto.UserProfileUpdate) error {
	_, err := s.db.Exec("DELETE FROM priority WHERE user_id = ?", id)
	if err != nil {
		return err
	}

	for i, priority := range up.PriorityOption {
		_, err := s.db.Exec("INSERT INTO priority (user_id, priority_id, rank) VALUES (?, ?, ?)", id, priority, i+1)
		if err != nil {
			return err
		}
	}

	_, err = s.db.Exec("UPDATE user_profile SET profile_image_url = ?, phone_number = ?, home_address = ?, card_id = ? WHERE id = ?", up.ProfileURL, up.PhoneNumber, up.HomeAddress, up.CardID, id)

	if err != nil {
		return err
	}

	return nil
}
