package controllers

import (
	"github.com/PrinceNarteh/go-auth/database"
	"github.com/PrinceNarteh/go-auth/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	hashedPassword, _ := models.HashPassword(data["password"])

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: hashedPassword,
	}

	database.DB.Create(&user)

	return c.JSON(data)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := models.ComparePassword(data["password"], user.Password); !err {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	return c.JSON(user)
}
