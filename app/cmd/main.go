package main

import (
	"log"
	"reservator/api/routes"
	"reservator/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db := utils.InitDBConnection()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.SetupShopRoutes(app, db)

	log.Fatal(app.Listen(":9999"))
}
