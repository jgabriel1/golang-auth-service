package repo

import (
	"database/sql"
	"errors"
	"golang-auth-service/src/models"

	"github.com/google/uuid"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	repo := UsersRepository{db}

	return &repo
}

func (repo *UsersRepository) All() ([]models.User, error) {
	rows, err := repo.db.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User

		if err := rows.Scan(user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo *UsersRepository) FindById(id uuid.UUID) (*models.User, error) {
	var user models.User

	result, err := repo.db.Query(
		"SELECT * FROM users WHERE id = $1;",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	result.Next()

	if err := result.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, errors.New("Couldn't find user!")
	}
	return &user, nil
}

func (repo *UsersRepository) FindByName(name string) (*models.User, error) {
	var user models.User

	result, err := repo.db.Query(
		"SELECT * FROM users WHERE username = $1;",
		name,
	)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	result.Next()

	if err := result.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, errors.New("Couldn't find user!")
	}
	return &user, nil
}

func (repo *UsersRepository) Create(username, password string) (*models.User, error) {
	var user models.User

	result, err := repo.db.Query(
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING *;",
		username,
		password,
	)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	result.Next()

	if err := result.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &user, nil
}
