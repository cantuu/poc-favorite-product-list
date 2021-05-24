package main

import (
	"list_product/app"
	"list_product/config"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.Start()

	app.Setup()
}
