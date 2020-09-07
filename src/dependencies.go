package server

import (
	"golang-auth-service/src/database"
	"golang-auth-service/src/repo"
	"golang-auth-service/src/routes"
	"golang-auth-service/src/services"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	c := dig.New()

	// Database connection
	c.Provide(database.NewDBConfig)
	c.Provide(database.NewDBConnection)

	// Repositories
	c.Provide(repo.NewUsersRepository)

	// Services
	c.Provide(services.NewRegisterUser)
	c.Provide(services.NewAuthenticateUser)
	c.Provide(services.NewAuthorizeUser)

	// Router
	c.Provide(routes.GetRouter)

	c.Provide(func() *ServerConfig {
		cfg := ServerConfig{Port: 8080}
		return &cfg
	})
	c.Provide(NewServer)

	return c
}
