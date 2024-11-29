package user

import (
	"medical-matching/constants/dto"

	"github.com/google/uuid"
)

func (s *UserService) AddPaymentMethod(userID int, card *dto.PaymentMethod) error {
	cardID := uuid.New().String()

	_, err := s.db.Exec("INSERT INTO payment_method (id, user_id, card_holder_name, card_number, exp_date, cvv) VALUES (?, ?, ?, ?, ?, ?)", cardID, userID, card.CardHolderName, card.CardNumber, card.ExpDate, card.Cvv)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetPaymentMethodList(userID int) ([]*dto.RetrievePaymentMethod, error) {
	rows, err := s.db.Query("SELECT id, card_number, exp_date FROM payment_method WHERE user_id = ?", userID)
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

func (s *UserService) DeletePaymentMethod(userID int, cardID string) error {
	return nil
}