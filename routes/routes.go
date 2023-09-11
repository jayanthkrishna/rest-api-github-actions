package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jayanthkrishna/rest-api-github-actions/contollers"
)

func GetRoutes(app *fiber.App) {

	app.Get("/", contollers.HelloWorld)

	app.Get("/api/v1/book", contollers.GetBooks)
	app.Get("/api/v1/book/:id", contollers.GetBook)
	app.Post("/api/v1/book", contollers.NewBook)
	app.Delete("/api/v1/book/:id", contollers.DeleteBook)
}
