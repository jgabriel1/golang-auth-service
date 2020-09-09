package main

import (
	"database/sql"
)

type CreateUsersTable01 struct {
	db *sql.DB
}

func (this *CreateUsersTable01) Up() (sql.Result, error) {
	return this.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id 			UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			username 	VARCHAR (64) UNIQUE NOT NULL,
			password 	VARCHAR (64) NOT NULL,
			created_at 	TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at 	TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`)
}

func (this *CreateUsersTable01) Down() (sql.Result, error) {
	return this.db.Exec(`
		DROP TABLE IF EXISTS users;
	`)
}
