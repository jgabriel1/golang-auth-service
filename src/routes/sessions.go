package routes

import (
	"encoding/json"
	"golang-auth-service/src/services"
	"golang-auth-service/src/utils"
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
			utils.JSONErrorResponse(w, err, http.StatusUnauthorized)
			return
		}

		js, err := json.Marshal(authenticatedData)
		if err != nil {
			utils.JSONErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	})

	return router
}
