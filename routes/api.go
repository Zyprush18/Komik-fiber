package routes

import (
	"github.com/Zyprush18/Komik-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func Api(c *fiber.App)  {
	c.Get("/api/users", controllers.ListUser)
	c.Post("/api/users/create", controllers.CreateUser)
	c.Get("/api/users/:id/edit", controllers.ShowUser)
	c.Put("/api/users/:id/update", controllers.UpdateUser)

}