package models

import "gorm.io/gorm"

// 회원(Member)
type Member struct {
	gorm.Model
	Email      string `gorm:"column:email" json:"email"`
	Password   string `json:"password"`
	FirstName  string `gorm:"column:first_name" json:"firstName"`
	LastName   string `gorm:"column:last_name" json:"lastName"`
	IsSignedIn bool   `gorm:"column:is_signed_in" json:"isSignedIn"`
}

func (Member) TableName() string {
	return "members"
}
