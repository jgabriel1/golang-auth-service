package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	createdAt time.Time
	updatedAt time.Time
}

// NewUser instantiates a new User object.
func NewUser(Username, Password string) *User {
	u := User{
		Id:        uuid.New(),
		Username:  Username,
		Password:  Password,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	return &u
}
