package main

import (
	"golang-auth-service/src/routes"
	"net/http"
)

func main() {
	router := routes.GetRouter()
	http.ListenAndServe(":8080", router)
}
