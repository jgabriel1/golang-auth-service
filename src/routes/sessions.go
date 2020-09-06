package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func sessionsRoutes(router *httprouter.Router) *httprouter.Router {
	router.POST("/login", func(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		fmt.Println("username: " + username)
		fmt.Println("password: " + password)

		w.WriteHeader(http.StatusAccepted)

		js, err := json.Marshal(struct {
			Username string
			Password string
		}{
			Username: username,
			Password: password,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Write(js)
	})

	return router
}
