package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// fiber app
	app := fiber.New()

	app.Use(cors.New())

	app.Static("/", "./client")

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "Users from backend",
		})
	})

	app.Listen(":4000")
}
