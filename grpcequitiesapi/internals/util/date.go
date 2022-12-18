package util

import (
	"fmt"
	"strings"
	"time"
)

var (
	errInvalidDateFormat   = "Please enter date in YYYY-MM-DD format"
	errDOBCannotBeInFuture = "Please enter a valid DOB"
	errInvalidDOB          = "Please enter a valid DOB"
	errInvalidAge          = "Please enter a valid DOB"

	errInvalidDate       = "Date validation failed %v"
	errInvalidDateLength = "Invalid date length %v"
	errInvalidInputDate  = "Please enter a valid date"
)

func parseDate(value string) (time.Time, error) {
	layout := time.RFC3339[:len(value)]
	return time.Parse(layout, value)
}

func ConvertDate(value string) (string, error) {
	if len(value) != 10 {
		date := fmt.Sprintf(errInvalidDate, value)
		return "", &BadRequest{ErrMessage: date}
	}
	seperator := value[4]
	dateArr := strings.Split(value, string(seperator))
	if len(dateArr) != 3 {
		date := fmt.Sprintf(errInvalidDateLength, value)
		return "", &BadRequest{ErrMessage: date}
	}
	newDateArr := []string{dateArr[0], dateArr[1], dateArr[2]}
	return strings.Join(newDateArr, "-"), nil
}

func checkIfFutureDate(dt time.Time) error {
	today := time.Now()
	if dt.After(today) {
		return &BadRequest{ErrMessage: errDOBCannotBeInFuture}
	}
	return nil
}
func ParseInputDate(dob string) (*time.Time, error) {
	dateStr, err := ConvertDate(dob)
	if err != nil {
		return nil, err
	}

	convertedDate, err := parseDate(dateStr + "T" + "00:00:00")
	if err != nil {
		return nil, &BadRequest{ErrMessage: errInvalidInputDate}
	}

	return &convertedDate, nil
}

func ParseDOB(dob string) (*time.Time, error) {
	dateStr, err := ConvertDate(dob)
	if err != nil {
		return nil, err
	}

	convertedDate, err := parseDate(dateStr + "T" + "00:00:00")
	if err != nil {
		return nil, &BadRequest{ErrMessage: errInvalidDOB}
	}

	//check if date is not in future
	err = checkIfFutureDate(convertedDate)
	if err != nil {
		return nil, err
	}

	yearDiff := calculateYearDiff(convertedDate, time.Now())
	if yearDiff > 200 {
		return nil, &BadRequest{ErrMessage: errInvalidAge}
	}
	return &convertedDate, nil
}

//reference https://golangr.com/difference-between-two-dates/
func calculateYearDiff(a, b time.Time) (year int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, _, _ := a.Date()
	y2, _, _ := b.Date()

	year = int(y2 - y1)

	return
}
