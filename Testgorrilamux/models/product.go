package models

type Product struct {
	Id              uint    `json:"id"`
	ProductName     string  `json:"product_name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	Image           string  `json:"image"`
	IsHot           bool    `json:"is_hot"`
	CategoryId      uint    `json:"category_id"`
	BrandId         uint    `json:"brand_id"`
	NumberAvailable uint    `json:"number_available"`
	CategoryName    string  `json:"category_name"`
	BrandName       string  `json:"brand_name"`
}