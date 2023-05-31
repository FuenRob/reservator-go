package utils

import (
	"log"
	"reservator/api/model"
	"reservator/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func createDBTables(db *gorm.DB) {
	db.AutoMigrate(&model.Booking{})
	db.AutoMigrate(&model.Shop{})
	db.AutoMigrate(&model.ShopHasDatesHasBooking{})
	db.AutoMigrate(&model.Dates{})
	db.AutoMigrate(&model.DatesHasBooking{})
}

func InitDBConnection() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.GetConnectionString()), &gorm.Config{})

	if err != nil {
		log.Panic("Error de conexión")
	} else {
		createDBTables(db)
	}

	return db
}
