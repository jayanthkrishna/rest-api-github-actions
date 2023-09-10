package contollers

import (
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, Welcome to the World my brother!!")
}
