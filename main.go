package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func getHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"input": `print("Enter some Lua code here and press submit in the toolbar.")`,
	})
}

func postHandler(c *fiber.Ctx) error {
	var output string

	input := c.FormValue("data")

	ctx, cancel := context.WithTimeout(c.Context(), 2*time.Second)
	defer cancel()

	output, err := executeLua(ctx, input)
	if err != nil {
		output = err.Error()
	}

	return c.Render("index", fiber.Map{
		"input":  input,
		"output": output,
	})
}

func main() {
	views := html.New("./views", ".tmpl")

	app := fiber.New(fiber.Config{
		Views: views,
	})

	app.
		Static("/", "./static").
		Get("/", getHandler).
		Post("/", postHandler)

	log.Fatal(app.Listen(":6969"))
}
