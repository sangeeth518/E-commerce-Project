package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string `json:"firstname" gorm:"not nul" validate:"required,min=2,max=50"`
	Lastname     string `json:"lastname" gorm:"not nul" validate:"required,min=1,max=50"`
	Email        string `json:"email"   gorm:"not null;unique"  validate:"email,required"`
	Password     string `json:"password" gorm:"not null"  validate:"required"`
	PhoneNumber  string `json:"phone"   gorm:"not null;unique" validate:"required"`
	Block_status bool   `JSON:"block_status" gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type Address struct {
	Id        uint   `json:"id" gorm:"unique;not null"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"-" gorm:"foreignkey:UserID"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Phone     string `json:"phone" gorm:"phone"`
	Pin       string `json:"pin" validate:"required"`
}

func (u *User) HashPassword(password string) (string, error) {

	byte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	u.Password = string(byte)
	return u.Password, nil

}
