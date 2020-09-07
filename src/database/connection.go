package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDBConnection(config *DBConfig) *sql.DB {
	psqlConfig := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.host,
		config.port,
		config.user,
		config.password,
		config.dbname,
	)

	db, err := sql.Open("postgres", psqlConfig)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	return db
}
