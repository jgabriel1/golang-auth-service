package service

import (
	"auth-service/domain/dtype"
	"auth-service/domain/model"
	"auth-service/domain/specification"
)

type RegisterUserService struct {
	repository    model.IUsersRepository
	nameIsUnique  specification.UserNameIsUnique
	emailIsUnique specification.EmailIsUnique
}

func NewRegisterUserService(repository *model.IUsersRepository) *RegisterUserService {
	service := &RegisterUserService{
		repository: *repository,
	}

	service.nameIsUnique = *specification.NewUserNameIsUnique(repository)
	service.emailIsUnique = *specification.NewEmailIsUnique(repository)

	return service
}

func (s *RegisterUserService) Execute(username, emailText, passwordText string) (*model.User, error) {
	var err error

	// Validate e-mail
	email, err := dtype.NewEmail(emailText)
	if err != nil {
		return nil, err
	}

	// Validate password
	password, err := dtype.NewPassword(passwordText)
	if err != nil {
		return nil, err
	}

	// TODO: Hash password
	passwordHash := dtype.NewPasswordHash([]byte(password.Value))

	// Check if user name is unique
	err = s.nameIsUnique.IsSatisfiedBy(username)
	if err != nil {
		return nil, err
	}

	// Check if email is unique
	err = s.emailIsUnique.IsSatisfiedBy(email)
	if err != nil {
		return nil, err
	}

	// Create user
	user, err := model.NewUser(username, email, passwordHash)
	if err != nil {
		return nil, err
	}

	// Persist created user
	err = s.repository.Register(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
