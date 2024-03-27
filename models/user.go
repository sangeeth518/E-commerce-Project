package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id           uint   `json:"id" gorm:"primarykey;unique"`
	FirstName    string `json:"fistname" gorm:"not nul" validate:"required,min=2,max=50"`
	Lastname     string `json:"lastname" gorm:"not nul" validate:"required,min=2,max=50"`
	Email        string `json:"email"   gorm:"not null;unique"  validate:"email,required"`
	Password     string `json:"password" gorm:"not null"  validate:"required"`
	PhoneNumber  int    `json:"phone"   gorm:"not null;unique" validate:"required"`
	Block_status bool   `JSON:"block_status" gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Adress struct {
	AdressId uint   `json:"adressid" gorm:"primarykey"`
	UserId   uint   `json:"userid" gorm:"not null" `
	Name     string `json:"name" gorm:"not null"`
	PhoneNum int    `json:"phonenum"`
	Pincode  int    `json:"pincode"`
	House    string `json:"house"`
	Area     string `json:"area"`
	Landmark string `json:"landmark"`
	City     string `json:"city"`
}
