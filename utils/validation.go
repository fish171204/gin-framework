package utils

import (
	"fmt"
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
