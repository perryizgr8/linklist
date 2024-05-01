package main

import (
	"linklist/links"
	"linklist/listing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	links.Init()
	engine := html.New("views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", listing.Index)
	app.Post("/add", listing.Add)
	app.Post("/edit", listing.Edit)
	app.Get("/delete/:id", listing.Delete)

	app.Listen(":3000")
}
