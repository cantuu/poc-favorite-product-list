package models

type ProductRequest struct {
	Price       string `json:"price" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Brand       string `json:"brand" binding:"required"`
	Title       string `json:"title" binding:"required"`
	ReviewScore string `json:"reviewScore" binding:"required"`
}
