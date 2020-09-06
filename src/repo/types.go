package repo

import (
	"golang-auth-service/src/models"
)

type IUsersRepository interface {
	Create(data CreateUserDTO) *models.User
	All() []*models.User
}

type CreateUserDTO struct {
	username, password string
}
