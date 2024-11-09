package routes

import (
	"github.com/Zyprush18/Komik-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func Web(c *fiber.App) {
	c.Get("/users", controllers.ShowUser)
	c.Post("/users/create", controllers.CreateUser)
}
