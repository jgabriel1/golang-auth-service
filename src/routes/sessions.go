package routes

import (
	"encoding/json"
	"golang-auth-service/src/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func sessionsRoutes(router *httprouter.Router, deps *RouteDependencies) *httprouter.Router {
	router.POST("/login", func(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		authenticateUser := services.NewAuthenticateUser(deps.UsersRepository)

		authenticatedData, err := authenticateUser.Execute(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		js, err := json.Marshal(authenticatedData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Write(js)
	})

	return router
}
