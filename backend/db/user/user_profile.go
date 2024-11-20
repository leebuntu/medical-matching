package user

import (
	"MedicalMatching/constants/dto"
	"database/sql"
	"fmt"
)

func (s *UserService) getEmailByID(id int) error {
	err := s.db.QueryRow("SELECT email FROM user WHERE id = ?", id).Scan(&s.profile.Email)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) getBasicInfoByID(id int) error {
	var profile_url sql.NullString
	var payment_id sql.NullString

	err := s.db.QueryRow("SELECT name, profile_image_url, phone_number, home_address, candy, payment_id FROM user_profile WHERE id = ?", id).Scan(&s.profile.Username, &profile_url, &s.profile.PhoneNumber, &s.profile.HomeAddress, &s.profile.Candy, &payment_id)
	if err != nil {
		return err
	}

	if profile_url.Valid {
		s.profile.ProfileURL = profile_url.String
	} else {
		s.profile.ProfileURL = ""
	}

	if payment_id.Valid {
		s.profile.PaymentMethod = "있는데 추후에 추가할꺼임"
	} else {
		s.profile.PaymentMethod = "없지롱"
	}

	return nil
}

func (s *UserService) getPriorArrByID(id int) error {
	s.profile.PriorityOption = make([]int, 5)
	var priorityID int

	err := s.db.QueryRow("SELECT priority_id FROM user_profile WHERE id = ?", id).Scan(&priorityID)
	if err != nil {
		return err
	}

	err = s.db.QueryRow("SELECT priority_1, priority_2, priority_3, priority_4, priority_5 FROM priority_set WHERE id = ?", priorityID).Scan(&s.profile.PriorityOption[0], &s.profile.PriorityOption[1], &s.profile.PriorityOption[2], &s.profile.PriorityOption[3], &s.profile.PriorityOption[4])
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserProfile(id int) (dto.UserProfile, error) {
	err := s.getEmailByID(id)
	if err != nil {
		return s.profile, err
	}

	err = s.getBasicInfoByID(id)
	if err != nil {
		return s.profile, err
	}

	err = s.getPriorArrByID(id)
	if err != nil {
		return s.profile, err
	}

	return s.profile, nil
}

func (s *UserService) findPrioritySetID(options []int) (int, error) {
	var id int

	err := s.db.QueryRow("SELECT id FROM priority_set WHERE priority_1 = ? AND priority_2 = ? AND priority_3 = ? AND priority_4 = ? AND priority_5 = ?", options[0], options[1], options[2], options[3], options[4]).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *UserService) UpdateUserProfile(id int, up dto.UserProfileUpdate) error {

	priorityID, err := s.findPrioritySetID(up.PriorityOption)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = s.db.Exec("UPDATE user_profile SET profile_image_url = ?, phone_number = ?, home_address = ?, payment_id = ?, priority_id = ? WHERE id = ?", up.ProfileURL, up.PhoneNumber, up.HomeAddress, up.PaymentID, priorityID, id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
