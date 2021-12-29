package main

import (
	"acp14/configs"
	"acp14/routes"
)

func main() {
	// Start Echo
	configs.Init()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8092"))
}
