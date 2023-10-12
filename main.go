package main

import (
	"TestBackEnd/handlers"
	"TestBackEnd/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	handlers.InitDataKabupaten()

	routes.KabupatenRoutes(app)

	// Start the server
	port := 3000
	log.Printf("Server started on :%d", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}