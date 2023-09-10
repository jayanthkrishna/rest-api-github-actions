package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jayanthkrishna/rest-api-github-actions/routes"
)

func main() {
	app := fiber.New()

	routes.GetRoutes(app)
	app.Listen(":3000")
}
