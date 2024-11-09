package main

import (
	"github.com/Zyprush18/Komik-fiber/databases"
	"github.com/Zyprush18/Komik-fiber/databases/migration"
	"github.com/Zyprush18/Komik-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func Server() {
	databases.ConnectDB()
	migration.RunMigrate()
	app := fiber.New()

	routes.Web(app)
	routes.Api(app)

	app.Listen(":8000")
}
