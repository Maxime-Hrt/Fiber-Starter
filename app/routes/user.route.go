package routes

import (
	"fiber-starter/app/controllers"
	"fiber-starter/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router) {
	api.Get("/me", middlewares.TokenMiddleware(), controllers.Me)
}
