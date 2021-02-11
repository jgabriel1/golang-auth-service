package dtype

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
)

type Password struct {
	Value string
}

func NewPassword(password string) (*Password, error) {
	p := &Password{
		Value: password,
	}

	err := p.IsValid()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid password: %s", err.Error()))
	}

	return p, nil
}

func (p *Password) IsValid() error {
	isAlphaNumeric := govalidator.IsAlphanumeric(p.Value)
	if !isAlphaNumeric {
		return errors.New("Password contains invalid characters.")
	}

	isLongerThanSixChars := govalidator.MinStringLength(p.Value, "6")
	if !isLongerThanSixChars {
		return errors.New("Password must be longer than 6 characters.")
	}

	containsOnlyLetters := govalidator.IsAlpha(p.Value)
	containsOnlyNumbers := govalidator.IsNumeric(p.Value)
	if containsOnlyLetters || containsOnlyNumbers {
		return errors.New("Password must contain letters and numbers.")
	}

	return nil
}

type PasswordHash struct {
	Value []byte
}

func NewPasswordHash(hash []byte) *PasswordHash {
	return &PasswordHash{
		Value: hash,
	}
}
