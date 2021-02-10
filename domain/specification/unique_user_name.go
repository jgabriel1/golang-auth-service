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

func (s *UserNameIsUnique) IsSatisfiedBy(username string) error {
	user, err := s.repository.FindByUserName(username)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("User name already taken.")
	}

	return nil
}
