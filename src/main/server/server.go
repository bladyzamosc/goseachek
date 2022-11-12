package server

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"goseachek/src/main/model"
)

var client ElasticClient

type Server struct {
}

func (Server) SetupServer() {
	client = ElasticClient{}.NewElasticClient()

	app := fiber.New()
	GetResultsEndpoint(app)
	IndexDataEndpoint(app)
	app.Listen(":3000")
}

func IndexDataEndpoint(app *fiber.App) fiber.Router {
	return app.Post("/index", func(ctx *fiber.Ctx) error {
		var requestIndex model.RequestIndex
		err := json.Unmarshal(ctx.Body(), &requestIndex)
		if err != nil {
			return ctx.SendString("ERROR, Data received, but the problem with parsing: " + err.Error())
		}
		requestIndex.PrintMe()
		return ctx.SendString("Data received")
	})
}

func GetResultsEndpoint(app *fiber.App) fiber.Router {
	return app.Get("/results", func(ctx *fiber.Ctx) error {
		queryValue := ctx.Query("text")
		return ctx.SendString("Text param: " + queryValue)
	})
}
