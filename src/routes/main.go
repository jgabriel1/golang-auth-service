package routes

import (
	"golang-auth-service/src/repo"
	"golang-auth-service/src/services"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/dig"
)

type RouteDependencies struct {
	dig.In

	// Repositories
	UsersRepository *repo.UsersRepository

	// Services
	RegisterUserService     *services.RegisterUser
	AuthenticateUserService *services.AuthenticateUser
	AuthorizeUserService    *services.AuthorizeUser
}

func GetRouter(deps RouteDependencies) *httprouter.Router {
	router := httprouter.New()

	usersRoutes(router, &deps)
	sessionsRoutes(router, &deps)

	return router
}
