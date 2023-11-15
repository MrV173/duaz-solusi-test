package main

import (
	"test-duaz-solusi/database"
	"test-duaz-solusi/pkg/mysql"
	"test-duaz-solusi/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	//Initializaing database
	mysql.DatabaseInit()

	//Initializing migration
	database.RunMigration()

    // Initialization fiber app
    app := fiber.New()

    // initializing route
	routes.RouteInit(app)
    // running on port 3000
    app.Listen(":3000")
}
