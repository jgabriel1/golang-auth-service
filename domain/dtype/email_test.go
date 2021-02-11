package dtype_test

import (
	"auth-service/domain/dtype"
	"testing"
)

func TestEmail(t *testing.T) {

	validEmailInput := "test@email.com"
	email, err := dtype.NewEmail(validEmailInput)

	if email == nil {
		t.Error("Passing a valid email should return an Email object.")
	}

	if err != nil {
		t.Error("Passing a valid email should not return any errors.")
	}

	invalidEmailInput := "invalid_email"
	_, err = dtype.NewEmail(invalidEmailInput)

	if err == nil {
		t.Error("Passing an invalid e-mail should return an error.")
	}
}
