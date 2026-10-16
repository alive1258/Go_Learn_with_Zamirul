package main

import (
	"errors"
	"fmt"
)

const (
	DefaultCurrency = "BDT"
	MaxTickets      = 10
)

var (
	ErrInvalidPrice    = errors.New("ticket price must be positive")
	ErrInvalidQuantity = errors.New("quantity must be between 1 and 10")
)

type Booking struct {
	UserName string
	Quantity int
	Price    int64 // Price per ticket in paisa.
}

func CalculateTotal(booking Booking) (int64, error) {
	if booking.Price <= 0 {
		return 0, ErrInvalidPrice
	}

	if booking.Quantity < 1 || booking.Quantity > MaxTickets {
		return 0, ErrInvalidQuantity
	}

	return booking.Price * int64(booking.Quantity), nil
}

func FormatBooking(booking Booking) (string, error) {
	total, err := CalculateTotal(booking)
	if err != nil {
		return "", fmt.Errorf("calculate booking total: %w", err)
	}

	summary := fmt.Sprintf(
		"User: %s | Quantity: %d | Total: %d paisa (%s)",
		booking.UserName,
		booking.Quantity,
		total,
		DefaultCurrency,
	)

	return summary, nil
}

func main() {
	booking := Booking{
		UserName: "Zamirul",
		Quantity: 2,
		Price:    50000,
	}

	summary, err := FormatBooking(booking)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(summary)
}
