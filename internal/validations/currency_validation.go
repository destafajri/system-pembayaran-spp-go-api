package validations

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	IDRRegex = `^IDR [0-9]+(\\.[0][0])?$` // IDR 1000.00
	USDRegex = `^USD [0-9]+(\\.[0][0])?$` // USD 5000.00
)

func IDRValidation(input string) (bool, error) {
	err := validation.Validate(input,
		validation.Required.Error("is required"),
		validation.Match(regexp.MustCompile(IDRRegex)).
			Error("must be IDR format example IDR 1000.00"),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func USDValidation(input string) (bool, error) {
	err := validation.Validate(input,
		validation.Required.Error("is required"),
		validation.Match(regexp.MustCompile(USDRegex)).
			Error("must be USD format example USD 5000.00"),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}
