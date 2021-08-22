package server

import (
	"github.com/gofiber/fiber/v2"
)

func Start(port string) {
	app := createApp()
	app.Listen(":" + port)
}

func createApp() *fiber.App {
	app := fiber.New()
	return app
}
