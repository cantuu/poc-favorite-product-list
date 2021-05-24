package app

import (
	"list_product/controllers"
)

var server = controllers.Server{}

func Setup() {
	server.Initialize()

	// seed.Load(server.DB)

	server.Run()

}
