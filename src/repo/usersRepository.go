package repo

import (
	"golang-auth-service/src/models"
)

type UsersRepository struct {
	data []*models.User
}

func NewUsersRepository() *UsersRepository {
	repo := UsersRepository{
		data: make([]*models.User, 0),
	}

	return &repo
}

func (repo *UsersRepository) All() []*models.User {
	return repo.data
}

func (repo *UsersRepository) Create(username, password string) *models.User {
	newUser := models.NewUser(
		username,
		password,
	)

	repo.data = append(repo.data, newUser)

	return newUser
}
