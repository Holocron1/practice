package myErrors

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Message string
	Field   string
}

func (s ValidationError) Error() string {
	return fmt.Sprintf("field %s: %s", s.Field, s.Message)
}

var ErrNotFound = errors.New("Field is empty")

func IsNotFound(err error) bool {
	if errors.Is(err, ErrNotFound) {
		return true
	} else {
		return false
	}
}

func AsValidation(err error) (*ValidationError, bool) {
	var v *ValidationError
	if errors.As(err, &v) {
		return v, true
	} else {
		return nil, false
	}
}
