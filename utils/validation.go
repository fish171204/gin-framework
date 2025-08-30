package utils

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

func ValidationRequired(fieldName, value string) error {
	if value == "" {
		return fmt.Errorf("%s is required", fieldName)
	}

	return nil
}

func ValidationStringLength(fieldName, value string, min, max int) error {
	l := len(value)
	if l < min || l > max {
		return fmt.Errorf("%s must be berween %d and %d characters", fieldName, min, max)
	}

	return nil
}

func ValidationRegex(value string, re *regexp.Regexp, errorMessage string) error {
	if !re.MatchString(value) {
		return fmt.Errorf("%s", errorMessage)
	}

	return nil
}

func ValidationPositiveInt(fieldName, value string) (int, error) {
	p, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be  a number", fieldName)
	}
	if p <= 0 {
		return 0, fmt.Errorf("%s must be positive", fieldName)
	}

	return p, nil
}

func ValidationUuid(fieldName, value string) (uuid.UUID, error) {
	uid, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s must be a UUID", fieldName)
	}

	return uid, nil
}
