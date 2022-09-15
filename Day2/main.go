package main

import (
	"github.com/alchemy/Day2/configs"
	"github.com/alchemy/Day2/routes"
)

func main() {

	configs.InitDB()

	app := routes.New()

	app.Start(":8080")
}
