package model

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey"`
	Reference  string    `json:"reference"`
	Name       string    `json:"name"`
	Time_open  time.Time `json:"time_open"`
	Time_close time.Time `json:"time_close"`
	Days_open  string    `json:"days_open"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ShopHasDatesHasBooking struct {
	ID                           uint `gorm:"primaryKey"`
	Shop_id                      int
	Dates_has_booking_dates_id   int
	Dates_has_booking_booking_id int
	CreatedAt                    time.Time
	UpdatedAt                    time.Time
}

func GetShops(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var shops []Shop
		db.Find(&shops)
		return c.JSON(shops)
	}
}

func GetShopById(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var shop Shop
		db.Find(&shop, id)
		return c.JSON(shop)
	}
}

func CreateShop(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		shop := new(Shop)

		if err := c.BodyParser(shop); err != nil {
			return c.Status(503).SendString(err.Error())
		}

		db.Create(&shop)
		return c.JSON(shop)
	}
}
