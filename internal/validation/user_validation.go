package validation

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func Phonenumber(input string) (bool, error) {
	if len([]rune(input)) < 9 {
		return false, errors.New("phone number length should be more than 9 digits")
	}
	
	regex := `^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`
	err := validation.Validate(input,
		validation.Required.Error("is required"),
		validation.Match(regexp.MustCompile(regex)).
		Error("must be a be phone number"),
	)

	if err != nil{
		return false , err
	}

	return true, nil
}