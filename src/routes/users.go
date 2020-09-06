package routes

import (
	"encoding/json"
	"golang-auth-service/src/repo"
	"golang-auth-service/src/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type usersPostJson struct {
	Username, Password string
}

func UsersRouter() *httprouter.Router {
	router := httprouter.New()
	usersRepository := repo.NewUsersRepository()

	router.POST("/users", func(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
		var data usersPostJson

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		}

		registerUser := services.RegisterUser{
			UserRepo: usersRepository,
		}

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