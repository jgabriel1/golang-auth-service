package services

import (
	"golang-auth-service/src/models"
)

type IRegisterUser interface {
	Execute(username, password string) *models.User
}
