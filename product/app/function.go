package app

import (
	"product/controllers"
	"product/seed"
)

var server = controllers.Server{}

// SetupRouter build gin engine
func Setup() {
	server.Initialize()

	seed.Load(server.DB)

	server.Run(":8080")

}
