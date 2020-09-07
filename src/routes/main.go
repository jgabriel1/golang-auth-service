package routes

import (
	"golang-auth-service/src/repo"

	"github.com/julienschmidt/httprouter"
)

type Dependencies struct {
	UsersRepository *repo.UsersRepository
}

func BuildDependencies(usersRepo *repo.UsersRepository) *Dependencies {
	d := Dependencies{
		UsersRepository: usersRepo,
	}

	return &d
}

func GetRouter(deps *Dependencies) *httprouter.Router {
	router := httprouter.New()

	usersRoutes(router, deps)
	sessionsRoutes(router, deps)

	return router
}
