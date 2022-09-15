package main

import (
	"Alterra/batch5/ORM/Part1/configs"
	"Alterra/batch5/ORM/Part1/routes"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// create a new echo instance

	configs.InitDB()

	app := routes.New()

	app.Start(":8080")
}
