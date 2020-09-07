package routes

import (
	"encoding/json"
	"golang-auth-service/src/services"
	"golang-auth-service/src/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type usersPostJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func usersRoutes(router *httprouter.Router, deps *RouteDependencies) *httprouter.Router {

	router.GET("/users/me", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		authorizeUser := services.NewAuthorizeUser(deps.UsersRepository)

		credentials, err := authorizeUser.Execute(&r.Header)
		if err != nil {
			utils.JSONErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		w.Write([]byte(credentials.UserID))
	})

	router.POST("/users", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var data usersPostJson

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		}

		registerUser := services.NewRegisterUser(deps.UsersRepository)

		user, err := registerUser.Execute(data.Username, data.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusCreated)

		js, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Write(js)
	})

	return router
}
