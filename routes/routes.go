package routes

import (
	"fiber-ebook-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Post("/upload", controllers.SaveEbooks)
	app.Get("/download", controllers.GetEbooks)
	//app.Get("/news", controllers.SaveNews)
	app.Get("/news/", controllers.GetNews)
	app.Get("/news-list/", controllers.FindPub)
}
