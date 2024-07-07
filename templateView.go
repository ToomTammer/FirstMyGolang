package main

import (
	"github.com/gofiber/fiber/v2"
)

func renderTemplate(c *fiber.Ctx) error {
  // Render the template.html with variable data
  return c.Render("template", fiber.Map{
    "Name": "World",
	"User": "ToomTam",
  })
}