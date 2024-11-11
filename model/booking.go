package model

import "time"

type Booking struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	CinemaID      string    `json:"cinema_id"`
	SeatID        string    `json:"seat_id"`
	Date          time.Time `json:"date"`
	Time          string    `json:"time"`
	PaymentMethod string    `json:"payment_method"`
}
