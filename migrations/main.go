package main

import (
	"database/sql"
	"fmt"

	server "golang-auth-service/src"
)

func parseCommandLineArgs() {

}

// TODO: iterate over types defined within this package to access each migration
// using package packages https://godoc.org/golang.org/x/tools/go/packages
func main() {
	container := server.BuildContainer()

	err := container.Invoke(func(db *sql.DB) {
		defer db.Close()

		migration := &CreateUsersTable01{db}

		res, err := migration.Up()
		if err != nil {
			panic(err)
		}

		fmt.Println(res)
	})

	if err != nil {
		panic(err)
	}
}
