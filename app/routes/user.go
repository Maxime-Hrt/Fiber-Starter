package routes

import (
	"fiber-starter/app/controllers"
	"fiber-starter/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router) {
	api.Get("/me", middlewares.AuthMiddleware(), controllers.Me)
	api.Delete("/delete-account", middlewares.AuthMiddleware(), controllers.DeleteUserById)
}
