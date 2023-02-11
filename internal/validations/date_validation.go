package validations

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	DateRegex = `\d{4}-\d{2}-\d{2}`
)

func DateValidation(input string) (bool, error) {
	err := validation.Validate(input,
		validation.Required.Error("is required"),
		validation.Match(regexp.MustCompile(DateRegex)).
			Error("must be date format example 2022-12-12"),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func YearValidation(input string) (bool, error) {
	year, err := strconv.Atoi(input)
	if err != nil {
		log.Println(err)
		return false, errors.New("invalid year format")
	}

	// Check if the year is within a valid range
	if year < 0 || year > 9999 {
		return false, errors.New("year must be between 0 and 9999")
	}

	// Check if the input string can be parsed as a valid time.Time value
	_, err = time.Parse("2006", input)
	if err != nil {
		log.Println(err)
		return false, errors.New("invalid year format")
	}

	return true, nil
}
