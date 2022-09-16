package main

import (
	"github.com/alchemy/Day3/configs"
	"github.com/alchemy/Day3/routes"
)

func main() {

	configs.InitDB()

	app := routes.New()

	app.Start(":8080")
}
