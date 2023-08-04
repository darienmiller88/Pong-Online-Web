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
        Views: html.New("./views", ".html"),
    })

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("route hit!")

		return c.JSON(map[string]interface{}{
			"message": "i ain't shit at coding :(",
		})
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"your id:": id,
		})
	})


	log.Fatal(app.Listen(":8080"))
}