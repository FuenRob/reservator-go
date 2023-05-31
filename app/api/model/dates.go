package model

import "time"

type Dates struct {
	ID        uint `gorm:"primaryKey"`
	Date      time.Time
	Service   string
	Busy      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DatesHasBooking struct {
	ID         uint `gorm:"primaryKey"`
	Dates_id   int
	Booking_id int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
