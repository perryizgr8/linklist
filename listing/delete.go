package listing

import (
	"linklist/links"

	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := links.Delete(id)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.Redirect("/")
}
