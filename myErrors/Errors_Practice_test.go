package myErrors

import (
	"testing"
)

func TestErrors(t *testing.T) {
	err := ErrNotFound
	got := IsNotFound(err)
	if got != true {
		t.Errorf("IsNotFound() = %v, want true", got)
	}

	var v *ValidationError
	error := &ValidationError{
		Field:   "email",
		Message: "invalid format",
	}
	v, ok := AsValidation(error)
	if !ok {
		t.Errorf("AsValidation() = %v, want ValidationError", v)
	}

	s := ValidationError{
		Message: "someMessage",
		Field:   "someField",
	}

	got2 := s.Error()
	target := "field someField: someMessage"
	if got2 != target {
		t.Errorf("waited for %v, got ", got2)
	}

}
