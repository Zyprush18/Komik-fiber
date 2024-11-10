package controllers

import (
	"log"

	"github.com/Zyprush18/Komik-fiber/databases"
	"github.com/Zyprush18/Komik-fiber/models/entity"
	"github.com/Zyprush18/Komik-fiber/models/requests"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ListComic(c *fiber.Ctx) error  {
	var comics []entity.Komik

	err := databases.DB.Find(&comics).Error
	
	if err != nil {
		log.Println(err)
	}

	if len(comics) == 0 {
		return c.JSON(fiber.Map{
			"message": "data belum ada",
		})
	}else{
		return c.JSON(comics)
	}

}


func CreateComics(c *fiber.Ctx) error {
	comic := new(requests.Komik)

	if err := c.BodyParser(comic); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(comic);err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "failed input user",
			"error": err.Error(),
		})
	}


	newcomic := entity.Komik{
		Name: comic.Name,
		Creator: comic.Creator,
		Publisher: comic.Publisher,
		Releases: comic.Releases,
		Image: comic.Image,
		Decsription: comic.Decsription,
	}

	if err := databases.DB.Create(&newcomic).Error;err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Failed Create Comics",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Succes Create new Comic ",
		"data": newcomic,
	})
}

func ShowComic(c *fiber.Ctx) error  {
	var comics []entity.Komik

	id := c.Params("id")

	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "unknown id",
		})
		return nil
	}

	if err := databases.DB.Where("id = ?",id).First(&comics).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "data not found",
		})
		return nil
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"data found",
		"data":comics,
	})
}


func UpdateComic(c *fiber.Ctx) error {
	comic := new(requests.Komik)
	var comics []entity.Komik

	id := c.Params("id")

	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "unknown id",
		})
		return nil
	}

	if err := c.BodyParser(comic);err != nil {
		return err
	}


	updatecomic := entity.Komik{
		Name: comic.Name,
		Creator: comic.Creator,
		Publisher: comic.Publisher,
		Releases: comic.Releases,
		Image: comic.Image,
		Decsription: comic.Decsription,
	}

	if err := databases.DB.Model(&comics).Where("id = ?", id).First(&comics).Updates(&updatecomic).Error;err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "failed update comic in id",
		})
		
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Update",
		"data": updatecomic,
	})
}


func DeleteComic(c *fiber.Ctx) error  {
	var comis []entity.Komik

	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unknown id",
		})
	}

	if err := databases.DB.Model(&comis).Where("id = ?",id).First(&comis).Delete(&comis).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed Delete comics in id",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Success Delete Comics",
	})
}