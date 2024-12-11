package utils

import (
	"medical-matching/constants"
	"regexp"
	"strings"
	"sync"
)

type FormatChecker struct {
	validPriorityOption map[int]bool
}

var once sync.Once
var instance *FormatChecker

func GetFormatChecker() *FormatChecker {
	once.Do(func() {
		m := make(map[int]bool)
		for i := 1; i <= constants.TotalPriority; i++ {
			m[i] = true
		}

		instance = &FormatChecker{
			validPriorityOption: m,
		}
	})

	return instance
}

func (f *FormatChecker) CheckEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func (f *FormatChecker) CheckAddress(address string, postalCode string) bool {
	//postalCodeRegex := regexp.MustCompile(`^\d{5,6}$`)

	// TODO: Implement address check with naver map api
	return true
}

func (f *FormatChecker) CheckCardData(cardNumber string, expDate string, cvv string) bool {
	cardNumberRegex := regexp.MustCompile(`^\d{16}$`)
	expDateRegex := regexp.MustCompile(`^\d{2}/\d{2}$`)
	cvvRegex := regexp.MustCompile(`^\d{3}$`)

	return cardNumberRegex.MatchString(strings.ReplaceAll(cardNumber, "-", "")) && expDateRegex.MatchString(expDate) && cvvRegex.MatchString(cvv)
}

func (f *FormatChecker) CheckURL(url string) bool {
	urlRegex := regexp.MustCompile(`^https?:\/\/[^ ]+$`)
	return urlRegex.MatchString(url)
}

func (f *FormatChecker) CheckPhoneNumber(phoneNumber string) bool {
	phoneNumberRegex := regexp.MustCompile(`^\d{10,11}$`)
	return phoneNumberRegex.MatchString(strings.ReplaceAll(phoneNumber, "-", ""))
}

func (f *FormatChecker) CheckPriorityOption(priorityOption []int) bool {
	seen := make(map[int]bool)

	if len(priorityOption) > 3 {
		return false
	}

	for _, num := range priorityOption {
		if !f.validPriorityOption[num] {
			return false
		}

		if seen[num] {
			return false
		}

		seen[num] = true
	}

	return true
}
