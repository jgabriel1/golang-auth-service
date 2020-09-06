package routes

import "github.com/julienschmidt/httprouter"

func GetRouter() *httprouter.Router {
	router := httprouter.New()

	usersRoutes(router)
	sessionsRoutes(router)

	return router
}
