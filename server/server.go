package server

import (
	"twit/funcs"

	"github.com/gofiber/fiber/v2"
)

func Start(port string) {
	app := createApp()
	app.Listen(":" + port)
}

func createApp() *fiber.App {
	app := fiber.New()
	app.Add("GET", "/func", funcs.Testing)
	return app
}
