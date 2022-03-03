package models

// 주문(Order)
type Order struct {
	Member
	Product
	MemberId      int `json:"memberId"`
	ProductId     int `json:"productId"`
	PurchasePrice int `json:"purchasePrice"`
	PurchaseDate  int `json:"PurchaseDate"`
}
