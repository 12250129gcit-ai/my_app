package main

import (
	_ "myapp/dataStore/postgres" // triggers init() – the underscore is important
	"myapp/routes"
)

func main() {
	routes.InitializeRoutes()
}
