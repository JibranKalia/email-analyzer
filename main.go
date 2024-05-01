package main

import (
	"email_analyzer/db"
	"email_analyzer/routes"
	// ...
)

func main() {
	db.InitDB()
	routes.RegisterRoutes()
	// Start web server
}
