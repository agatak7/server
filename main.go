package main

import (
	"go-trans/configs"
	"go-trans/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title User API documentation
// @version 1.0.0

// @host localhost:8080
// @BasePath /

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Welcome to transportation"})
	})

	app.Use(cors.New())

	//run database
	configs.ConnectDB()

	//routes
	routes.TransportRoute(app)

	//listen on port
	app.Listen(":8080")
}
