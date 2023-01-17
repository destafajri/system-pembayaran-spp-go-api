package validations

import (
	"errors"
	"regexp"
	"strings"

	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	UsernameRegex    = `^[a-zA-Z0-9]*[-]?[a-zA-Z0-9]*$`
	EmailRegex       = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	SpecialCharRegex = `([!@#$%^&*.?-])+`
	PhoneNumberRegex = `^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`
)

func PhonenumberValidation(input string) (bool, error) {
	if len([]rune(input)) < 9 {
		return false, errors.New("phone number length should be more than 9 digits")
	}

	err := validation.Validate(input,
		validation.Required.Error("is required"),
		validation.Match(regexp.MustCompile(PhoneNumberRegex)).
			Error("must be a phone number"),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func UsernameValidation(input string) (bool, error) {
	space := strings.Contains(input, " ")
	if space {
		return false, &exception.ValidationError{
			Message: "Password should not containt white space",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	if len([]rune(input)) < 4 || len([]rune(input)) > 30 {
		return false, errors.New("username length should be more than 4 digits and less than 30")
	}

	err := validation.Validate(input,
		validation.Required.Error("is required"),
		validation.Match(regexp.MustCompile(UsernameRegex)).
			Error("must be a username"),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func EmailValidation(input string) (bool, error) {
	if len([]rune(input)) < 4 || len([]rune(input)) > 40 {
		return false, errors.New("email length should be more than 4 digits and less than 40")
	}

	err := validation.Validate(input,
		validation.Required.Error("is required"),
		validation.Match(regexp.MustCompile(EmailRegex)).
			Error("must be a email"),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func PasswordValidation(password string) (bool, error) {
	space := strings.Contains(password, " ")
	if space {
		return false, &exception.ValidationError{
			Message: "Password should not containt white space",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	if len(password) < 6 {
		return false, &exception.ValidationError{
			Message: "Password should be of 6 characters long",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	done, err := regexp.MatchString("([a-z])+", password)
	if err != nil {
		return false, err
	}

	if !done {
		return false, &exception.ValidationError{
			Message: "Password should contain atleast one lower case character",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	done, err = regexp.MatchString("([A-Z])+", password)
	if err != nil {
		return false, err
	}

	if !done {
		return false, &exception.ValidationError{
			Message: "Password should contain atleast one upper case character",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	done, err = regexp.MatchString("([0-9])+", password)
	if err != nil {
		return false, err
	}

	if !done {
		return false, &exception.ValidationError{
			Message: "Password should contain atleast one number",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	return true, nil
}
