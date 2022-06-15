package main

import (
	"fiber-ebook-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 40 * 1024 * 1024, // this is the default limit of 40MB
	})

	routes.Router(app)
	err := app.Listen(":80")
	if err != nil {
		log.Fatal(err)
	}
}
