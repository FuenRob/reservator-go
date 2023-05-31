package routes

import (
	"reservator/api/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupShopRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/v1/shop", model.GetShops(db))
	app.Get("/api/v1/shop/:id", model.GetShopById(db))
	app.Post("/api/v1/shop", model.CreateShop(db))
	//app.Put("/api/v1/shop", func(c *fiber.Ctx) error {})
}
