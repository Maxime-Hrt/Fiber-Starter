package controllers

import (
	"fiber-starter/app/models"
	"fiber-starter/app/services"

	"github.com/gofiber/fiber/v2"
)

func Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func DeleteUserById(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	if err := services.DeleteUserById(user.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
