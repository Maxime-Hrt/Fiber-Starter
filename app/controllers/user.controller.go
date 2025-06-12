package controllers

import (
	"fiber-starter/app/models"

	"github.com/gofiber/fiber/v2"
)

func Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}
