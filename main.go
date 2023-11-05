package main

import (
	"SkripsiBebek/db"
	"SkripsiBebek/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":38560"))
}
