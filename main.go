package main

import (
	"github.com/felipeospina21/todo_fiber_api/database"
	"github.com/felipeospina21/todo_fiber_api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	api.Get("/todo", routes.GetAll)
	api.Post("/todo", routes.Create)
	api.Put("/todo/:id", routes.Update)
	api.Delete("/todo/:id", routes.Delete)

	app.Listen(":8080")
}
