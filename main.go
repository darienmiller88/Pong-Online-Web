package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main()  {
	fmt.Println("Hello world")

	app := fiber.New(fiber.Config{
        Views: html.New("./client", ".html"),
		Prefork: true,
    })

	app.Static("/", "./client")
	
	app.Get("/pong", func(c *fiber.Ctx) error {
		return c.Render("index", "Title")
	})

	app.Get("/api/", func(c *fiber.Ctx) error {
		fmt.Println("route hit!")

		return c.JSON(fiber.Map{
			"message": "i ain't shit at coding :(",
		})
	})

	app.Get("/api/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"your id:": id,
		})
	})


	log.Fatal(app.Listen(":8080"))
}