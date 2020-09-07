package main

import (
	server "golang-auth-service/src"
)

func main() {
	container := server.BuildContainer()

	err := container.Invoke(func(sv *server.Server) {
		sv.Run()
	})

	if err != nil {
		panic(err)
	}
}
