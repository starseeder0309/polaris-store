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
