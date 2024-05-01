package listing

import (
	"linklist/links"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "linklist",
		"Links": links.Get(),
	})
}
