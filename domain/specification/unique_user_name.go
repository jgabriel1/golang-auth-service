package specification

import (
	"auth-service/domain/model"
	"errors"
)

type UserNameIsUnique struct {
	repository model.IUsersRepository
}

func NewUserNameIsUnique(repository *model.IUsersRepository) *UserNameIsUnique {
	return &UserNameIsUnique{
		repository: *repository,
	}
}

func (s *UserNameIsUnique) IsSatisfiedBy(user *model.User) error {
	user, err := s.repository.FindByUserName(user.Name)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("User name already taken.")
	}

	return nil
}
