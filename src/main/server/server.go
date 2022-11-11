package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
}

func (Server) SetupServer() {
	app := fiber.New()

	GetResults(app)
	IndexData(app)

	app.Listen(":3000")
}

func IndexData(app *fiber.App) fiber.Router {
	return app.Post("/hello", func(ctx *fiber.Ctx) error {
		var _body = ctx.Body()
		fmt.Println(string(_body))
		return ctx.SendString("Data received")
	})
}

func GetResults(app *fiber.App) fiber.Router {
	return app.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("I am running")
	})
}
