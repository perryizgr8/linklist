package listing

import (
	"linklist/links"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":  viper.GetString("app_name"),
		"IdName": viper.GetString("id_name"),
		"Links":  links.Get(),
	})
}
