package main

import (
	"POSSystem/initializers"
	"POSSystem/routes"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectDB()
}
func main() {

	r := routes.SetupRoutes()
	r.Run()
}
