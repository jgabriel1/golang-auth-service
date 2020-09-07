package server

import (
	"golang-auth-service/src/database"
	"golang-auth-service/src/repo"
	"golang-auth-service/src/routes"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	c := dig.New()

	// Database connection
	c.Provide(database.NewDBConfig)
	c.Provide(database.NewDBConnection)

	// Repositories
	c.Provide(repo.NewUsersRepository)

	// Router
	c.Provide(routes.GetRouter)

	c.Provide(func() *ServerConfig {
		cfg := ServerConfig{Port: 8080}
		return &cfg
	})
	c.Provide(NewServer)

	return c
}
