package models

// 상품(Product)
type Product struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	PromotedPrice float64 `json:"PromotedPrice"`
	Image         string  `json:"image"`
	ImageAlt      string  `json:"imageAlt"`
}
