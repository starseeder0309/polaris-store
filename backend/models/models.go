package models

// 회원(Member)
type Member struct {
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	IsSignedIn bool    `json:"isSignedIn"`
	Orders     []Order `json:"orders"`
}

// 상품(Product)
type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"`
	Image       string  `json:"image"`
	ImageAlt    string  `json:"imageAlt"`
}

// 주문(Order)
type Order struct {
	Member
	Product
	MemberId      int `json:"memberId"`
	ProductId     int `json:"productId"`
	PurchasePrice int `json:"purchasePrice"`
	PurchaseDate  int `json:"PurchaseDate"`
}
