package main

import (
	"github.com/gofiber/fiber"
	"github.com/kumarraja/go-tutorials/tree/master/intermediate/go-fiber-tutorial/book"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)

}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
