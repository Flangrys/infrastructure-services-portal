package app

import (
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {

	app := fiber.New(fiber.Config{
		AppName:      "ISP",
		ServerHeader: "ISP/",
	})

	return app
}
