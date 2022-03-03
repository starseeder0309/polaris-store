package models

import (
	"time"

	"gorm.io/gorm"
)

// 주문(Order)
type Order struct {
	gorm.Model
	Member
	Product
	MemberId      int       `gorm:"column:member_id"`
	ProductId     int       `gorm:"column:product_id"`
	PurchasePrice int       `gorm:"column:purchase_price" json:"purchasePrice"`
	PurchaseDate  time.Time `gorm:"column:purchase_date" json:"PurchaseDate"`
}

func (Order) TableName() string {
	return "orders"
}
