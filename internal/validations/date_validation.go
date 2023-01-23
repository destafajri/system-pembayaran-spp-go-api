package validations

import (
	"regexp"

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