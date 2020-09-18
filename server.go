package main

import "github.com/jamestjw/coup-vin/controllers"

var server = controllers.Server{}

func runServer() {
	server.Initialize("production")
	server.Run(":8080")
}
