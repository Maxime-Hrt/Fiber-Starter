package main

import (
	"context"
	"fiber-starter/app/db"
	"fiber-starter/app/routes"
	"fiber-starter/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.LoadEnv(".env")
	db.ConnectDB()

	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173", // Set specific origin instead of wildcard
		AllowHeaders:     "Content-Type, Authorization",
	}))

	routes.SetupRoutes(app)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Start server
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatal("Error starting server: ", err)
		}
	}()

	log.Println("Server started on port 3000")

	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatal("Error shutting down server: ", err)
	}

	db.CloseDB()

	log.Println("Server stopped")
}
