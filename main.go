package main

import (
	"demo-gofiber/mockup"
	"demo-gofiber/query"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", HelloHandler)

	mockRepository := mockup.NewMockRepository()
	queryHandler := query.NewHandler(mockRepository)

	v1 := app.Group("/v1")
	v1.Get("/all", queryHandler.GetAllHandler)
	v1.Get("/sum", queryHandler.SumHandler)
	app.Listen(":3000")
}

func HelloHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello, World!",
	})
}
