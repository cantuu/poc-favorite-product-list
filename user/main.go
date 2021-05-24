package main

import (
	"user/app"
	"user/config"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.Start()

	app.Setup()
}
