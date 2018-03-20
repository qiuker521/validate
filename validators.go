package validate

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func validateNotNull(path string, value interface{}) error {
	if value != nil {
		return nil
	} else {
		return errors.New("Value of " + path + " should not be nil")
	}
}

func validateNoSpace(path string, value interface{}) error {
	err, v := ConvertToString(value)
	if err != nil {
		return err
	}
	if strings.ContainsAny(v, " 	") {
		return errors.New("Value of " + path + " should not contains space")
	}
	return nil
}

func validateNotEmpty(path string, value interface{}) error {
	err, v := ConvertToString(value)
	if err != nil {
		return err
	}
	if strings.Trim(strings.Trim(v, "	"), " ") == "" {
		return errors.New("Value of " + path + " should not be empty")
	}
	return nil
}

func validateEmail(path string, value interface{}) error {
	err, v := ConvertToString(value)
	if err != nil {
		return err
	}
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if Re.MatchString(v) {
		return nil
	} else {
		return errors.New("String of " + path + " must be an email, but none email string privided")
	}
}

func ConvertToString(s interface{}) (error, string) {
	if v, ok := s.(string); ok {
		return nil, v
	} else {
		return errors.New(fmt.Sprintf("Value of %#v is not an email.", s)), ""
	}
}
