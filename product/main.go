package main

import (
	"product/app"
	"product/config"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.Start()

	app.Setup()
}
