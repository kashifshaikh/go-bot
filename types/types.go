package types

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Parses expiry date in MM/YYYY format and verifies it has not expired yet
func ExpValidate(expDate string) (month, year int, err error) {
	now := time.Now()
	exp := strings.Split(expDate, "/")

	if len(exp) != 2 || len(exp[0]) != 2 || len(exp[1]) != 4 {
		return 0, 0, errors.New("invalid format: " + expDate)
	}
	year, err = strconv.Atoi(exp[1])
	if err != nil {
		return 0, 0, errors.New("invalid year: " + exp[1])
	}
	month, err = strconv.Atoi(exp[0])
	if err != nil || month < 1 || month > 12 {
		return 0, 0, errors.New("invalid month: " + exp[0])
	}
	if year < now.Year() || (year == now.Year() && month < int(now.Month())) {
		return 0, 0, errors.New("has expired: " + exp[0] + "/" + exp[1])
	}
	return month, year, nil
}

func GetExpMonthYear(expDate string) (month, year string, err error) {
	exp := strings.Split(expDate, "/")

	if len(exp) != 2 || len(exp[0]) != 2 || len(exp[1]) != 4 {
		return "", "", errors.New("invalid format: " + expDate)
	}
	return month, year, nil
}
