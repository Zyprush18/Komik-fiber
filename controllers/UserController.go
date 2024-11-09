package controllers

import (
	"log"

	"github.com/Zyprush18/Komik-fiber/databases"
	"github.com/Zyprush18/Komik-fiber/models/entity"
	"github.com/Zyprush18/Komik-fiber/models/requests"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


func ShowUser(c *fiber.Ctx) error  {
	var users []entity.User

	err := databases.DB.Find(&users).Error

	if err != nil {
		log.Println(err)	
	}

	return c.JSON(users)
}


func CreateUser(c *fiber.Ctx) error {
	user := new(requests.UserRequests)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	// validasi sebelum menyimpan data
	validation := validator.New()
	if err := validation.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed Input User",
			"errors": err.Error(),
		})
	}




	newUser := entity.User{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}

	if err := databases.DB.Create(&newUser).Error;err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Failed Create User",
			})
	}

	return c.JSON(fiber.Map{
		"message": "Succes Create User",
		"data": newUser,
	})

}