package specification

import (
	"auth-service/domain/model"
	"errors"
)

type UserEmailIsUnique struct {
	repository model.IUsersRepository
}

func NewUserEmailIsUnique(repository *model.IUsersRepository) *UserEmailIsUnique {
	return &UserEmailIsUnique{
		repository: *repository,
	}
}

func (s *UserEmailIsUnique) IsSatisfiedBy(user *model.User) error {
	user, err := s.repository.FindByEmail(&user.Email)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("This e-mail is already taken.")
	}

	return nil
}
