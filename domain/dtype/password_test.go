package dtype_test

import (
	"auth-service/domain/dtype"
	"testing"
)

func TestPassword(t *testing.T) {

	validPassword := "password123"
	password, err := dtype.NewPassword(validPassword)

	if password == nil {
		t.Error("Passing a valid password should return a Password object.")
	}

	if err != nil {
		t.Error("Passing a valid password should not return any errors.")
	}

	tooShortPassword := "pwd1"
	_, err = dtype.NewPassword(tooShortPassword)

	if err == nil {
		t.Error("Passing a password shorter than 6 characters should return an error.")
	}

	noNumbersPassword := "longPassword"
	_, err = dtype.NewPassword(noNumbersPassword)

	if err == nil {
		t.Error("Passing a password without any number characters should return an error.")
	}

	onlyNumbersPassword := "123456"
	_, err = dtype.NewPassword(onlyNumbersPassword)

	if err == nil {
		t.Error("Passing a password without any letters in it should return an error")
	}
}

func TestPasswordHash(t *testing.T) {
	anyHash := []byte("anyPassword")

	passwordHash := dtype.NewPasswordHash(anyHash)

	if passwordHash == nil {
		t.Error("Passing any hash should return a password hash object, there are no validation steps here.")
	}
}
