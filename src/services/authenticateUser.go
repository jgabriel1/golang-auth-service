package services

import (
	"errors"
	"golang-auth-service/src/security"

	"golang-auth-service/src/models"
	"golang-auth-service/src/repo"

	"golang.org/x/crypto/bcrypt"
)

type AuthenticateUser struct {
	usersRepo *repo.UsersRepository
}

type AuthenticateUserResponse struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}

func NewAuthenticateUser(usersRepo *repo.UsersRepository) *AuthenticateUser {
	this := AuthenticateUser{
		usersRepo: usersRepo,
	}

	return &this
}

func (this *AuthenticateUser) compareHashAndPassword(hash, password string) error {
	hashBytes := []byte(hash)
	passwordBytes := []byte(password)

	return bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)
}

func (this *AuthenticateUser) Execute(username, password string) (*AuthenticateUserResponse, error) {
	user, err := this.usersRepo.FindByName(username)
	if err != nil {
		return nil, errors.New("Wrong username/password combination.")
	}

	if err := this.compareHashAndPassword(user.Password, password); err != nil {
		return nil, errors.New("Wrong username/password combination.")
	}

	accessToken, err := security.GenerateSignedJWT(user.Id.String())

	response := &AuthenticateUserResponse{
		User:  user,
		Token: accessToken,
	}

	return response, nil
}
