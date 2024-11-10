package routes

import (
	"github.com/Zyprush18/Komik-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func Api(c *fiber.App)  {
	// api user
	c.Get("/api/users", controllers.ListUser)
	c.Post("/api/users/create", controllers.CreateUser)
	c.Get("/api/users/:id/show", controllers.ShowUser)
	c.Put("/api/users/:id/update", controllers.UpdateUser)
	c.Delete("/api/users/:id/delete", controllers.DeleteUser)


	// api comic
	c.Get("/api/comics",controllers.ListComic)
	c.Post("/api/comics/create",controllers.CreateComics)
	c.Get("/api/comics/:id/show",controllers.ShowComic)
	c.Put("/api/comics/:id/update",controllers.UpdateComic)
	c.Delete("/api/comics/:id/delete",controllers.DeleteComic)
}