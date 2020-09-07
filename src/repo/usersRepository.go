package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-auth-service/src/models"

	"github.com/google/uuid"
)

type UsersRepository struct {
	db   *sql.DB
	data []*models.User
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	repo := UsersRepository{
		db:   db,
		data: make([]*models.User, 0),
	}

	return &repo
}

func (repo *UsersRepository) All() []*models.User {
	return repo.data
}

func (repo *UsersRepository) FindById(id uuid.UUID) (*models.User, error) {
	for _, user := range repo.data {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, errors.New("User does not exist.")
}

func (repo *UsersRepository) FindByName(name string) (*models.User, error) {
	for _, user := range repo.data {
		if user.Username == name {
			return user, nil
		}
	}

	return nil, errors.New("User does not exist.")
}

func (repo *UsersRepository) Create(username, password string) (*models.User, error) {
	newUser := models.NewUser(
		username,
		password,
	)

	repo.data = append(repo.data, newUser)

	fmt.Println(repo.data)

	return newUser, nil
}
