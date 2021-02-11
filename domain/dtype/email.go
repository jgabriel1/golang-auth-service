package dtype

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

type Email struct {
	Value string
}

func NewEmail(email string) (*Email, error) {
	e := &Email{
		Value: email,
	}

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Email) IsValid() error {
	isValidEmail := govalidator.IsEmail(e.Value)
	if !isValidEmail {
		return errors.New("Invalid e-mail.")
	}

	return nil
}
