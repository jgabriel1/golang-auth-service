package services

import (
	"errors"
	"golang-auth-service/src/models"
	"golang-auth-service/src/repo"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	UserRepo *repo.UsersRepository
}

func (this *RegisterUser) hashPassword(password string) (string, error) {
	pwdBytes := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(
		pwdBytes,
		bcrypt.MinCost,
	)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (this *RegisterUser) Execute(username, password string) (*models.User, error) {
	hashedPwd, err := this.hashPassword(password)

	if err != nil {
		return nil, errors.New("Error hashing password:" + err.Error())
	}

	createdUser, err := this.UserRepo.Create(username, string(hashedPwd))

	if err != nil {
		return nil, errors.New("Error creating user:" + err.Error())
	}

	return createdUser, nil
}
