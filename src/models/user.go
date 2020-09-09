package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser instantiates a new User object.
func NewUser(Username, Password string) *User {
	u := User{
		Id:        uuid.New(),
		Username:  Username,
		Password:  Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &u
}
