package main

type Migration interface {
	Up()
	Down()
}
