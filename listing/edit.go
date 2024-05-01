package listing

import (
	"fmt"
	"linklist/links"

	"github.com/gofiber/fiber/v2"
)

func Edit(c *fiber.Ctx) error {
	id := c.FormValue("id")
	nick := c.FormValue("nick")
	url := c.FormValue("url")
	fmt.Println("edit:", id, nick, url)
	err := links.Edit(id, nick, url)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.Redirect("/")
}
