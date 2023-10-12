package routes

import (
	"TestBackEnd/handlers"

	"github.com/gofiber/fiber/v2"
)

func KabupatenRoutes(app *fiber.App) {
	kabupaten := app.Group("/kabupaten")
	kabupaten.Get("/list", handlers.ListKabupaten)
	kabupaten.Post("/detail_kabupaten", handlers.DetailKabupatenData)
}