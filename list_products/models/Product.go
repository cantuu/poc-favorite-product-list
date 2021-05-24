package models

import "time"

type Product struct {
	ID          uint32    `json:"id"`
	Price       string    ` json:"price"`
	Image       string    `json:"image"`
	Brand       string    `json:"brand"`
	Title       string    `json:"title"`
	ReviewScore string    `json:"reviewScore"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
