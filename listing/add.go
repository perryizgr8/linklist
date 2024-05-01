package listing

import (
	"fmt"
	"linklist/links"

	"github.com/gofiber/fiber/v2"
)

func Add(c *fiber.Ctx) error {
	nick := c.FormValue("nick")
	url := c.FormValue("url")
	fmt.Println("add:", nick, url)
	err := links.Add(nick, url)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.Redirect("/")
}
