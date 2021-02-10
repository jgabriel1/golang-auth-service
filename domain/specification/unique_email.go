package specification

import (
	"auth-service/domain/dtype"
	"auth-service/domain/model"
	"errors"
)

type EmailIsUnique struct {
	repository model.IUsersRepository
}

func NewEmailIsUnique(repository *model.IUsersRepository) *EmailIsUnique {
	return &EmailIsUnique{
		repository: *repository,
	}
}

func (s *EmailIsUnique) IsSatisfiedBy(email *dtype.Email) error {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("This e-mail is already taken.")
	}

	return nil
}
