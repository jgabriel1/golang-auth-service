package model

import (
	"auth-service/domain/dtype"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID          `json:"id"`
	Name         string             `json:"name"`
	Email        dtype.Email        `json:"email"`
	PasswordHash dtype.PasswordHash `json:"-"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}

type IUsersRepository interface {
	Register(user *User) error
	FindByEmail(email *dtype.Email) (*User, error)
	FindByUserName(username string) (*User, error)
}

func NewUser(name string, email *dtype.Email, passwordHash *dtype.PasswordHash) (*User, error) {
	user := &User{
		Name:         name,
		Email:        *email,
		PasswordHash: *passwordHash,
	}

	user.Id = uuid.New()
	user.CreatedAt = time.Now()

	err := user.IsValid()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) IsValid() error {
	return nil
}
