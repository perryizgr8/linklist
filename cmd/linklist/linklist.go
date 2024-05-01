package main

import (
	"linklist/follow"
	"linklist/links"
	"linklist/listing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	links.Init()
	engine := html.New("views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", listing.Index)
	app.Get("/:nick", follow.RedirectByNick)
	id_name := viper.GetString("id_name")
	app.Get("/"+id_name+"/:id", follow.RedirectById)
	app.Post("/add", listing.Add)
	app.Post("/edit", listing.Edit)
	app.Get("/delete/:id", listing.Delete)

	app.Listen(":3000")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// overrides from env
	viper.SetEnvPrefix("linklist")
	viper.AutomaticEnv()
}
