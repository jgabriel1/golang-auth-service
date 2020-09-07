package routes

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type usersPostJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func usersRoutes(router *httprouter.Router, deps *RouteDependencies) *httprouter.Router {

	router.GET("/users/me", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		credentials, err := deps.AuthorizeUserService.Execute(&r.Header)
		if err != nil {
			JSONErrorResponse(w, err, http.StatusUnauthorized)
			return
		}

		userId, _ := uuid.Parse(credentials.UserID)

		user, err := deps.UsersRepository.FindById(userId)
		if err != nil {
			JSONErrorResponse(w, err, http.StatusNotFound)
		}

		js, _ := json.Marshal(user)

		w.Write(js)
	})

	router.POST("/users", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var data usersPostJson

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		}

		user, err := deps.RegisterUserService.Execute(data.Username, data.Password)
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
