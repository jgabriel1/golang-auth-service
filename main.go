package main

import (
	"golang-auth-service/src/routes"
	"net/http"
)

func main() {
	usersRouter := routes.UsersRouter()
	http.ListenAndServe(":8080", usersRouter)
}
