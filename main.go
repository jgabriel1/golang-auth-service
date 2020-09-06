package main

import (
	"encoding/json"
	"golang-auth-service/src/repo"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	usersRepo := repo.NewUsersRepository()

	usersRepo.Create("gabriel", "1234")
	usersRepo.Create("josé", "1234")

	users := usersRepo.All()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
