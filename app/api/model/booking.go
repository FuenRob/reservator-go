package model

import "time"

type Booking struct {
	ID        uint `gorm:"primaryKey"`
	Shop_id   int
	Name      string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
