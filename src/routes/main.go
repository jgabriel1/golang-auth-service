package routes

import (
	"golang-auth-service/src/repo"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/dig"
)

type RouteDependencies struct {
	dig.In

	UsersRepository *repo.UsersRepository
}

func GetRouter(deps RouteDependencies) *httprouter.Router {
	router := httprouter.New()

	usersRoutes(router, &deps)
	sessionsRoutes(router, &deps)

	return router
}
