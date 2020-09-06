package database

import (
	"os"
	"strconv"
)

type DBConfig struct {
	user     string
	password string
	dbname   string
	host     string
	port     int64
}

func NewDBConfig() *DBConfig {
	this := DBConfig{
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   os.Getenv("DB_DBNAME"),
		host:     os.Getenv("DB_HOST"),
		port: func() int64 {
			parsedPort, _ := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 32)
			return parsedPort
		}(),
	}

	return &this
}
