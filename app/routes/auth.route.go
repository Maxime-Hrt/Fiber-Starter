package routes

import (
	"fiber-starter/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router) {
	api.Post("/signup", controllers.SignUp)
}
