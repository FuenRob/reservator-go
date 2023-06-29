package main

import (
	"log"
	"reservator/api/routes"
	"reservator/utils"

	"github.com/gofiber/fiber/v2"
)

func newApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	db := utils.InitDBConnection()

	routes.SetupShopRoutes(app, db)

	return app
}

func main() {
	log.Fatal(newApp().Listen(":9999"))
}
