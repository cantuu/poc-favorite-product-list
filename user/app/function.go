package app

import (
	"user/controllers"
)

var server = controllers.Server{}

func Setup() {
	server.Initialize()

	server.Run()

}
