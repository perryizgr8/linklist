package follow

import (
	"linklist/links"

	"github.com/gofiber/fiber/v2"
)

func RedirectByNick(c *fiber.Ctx) error {
	nick := c.Params("nick")
	url, err := links.GetUrlByNick(nick)
	if url == "" || err != nil {
		return c.Redirect("/")
	}
	return c.Redirect(url)
}

func RedirectById(c *fiber.Ctx) error {
	id := c.Params("id")
	url, err := links.GetUrlById(id)
	if url == "" || err != nil {
		return c.Redirect("/")
	}
	return c.Redirect(url)
}
