package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jayanthkrishna/rest-api-github-actions/contollers"
)

func GetRoutes(app *fiber.App) {

	app.Get("/", contollers.HelloWorld)

}
