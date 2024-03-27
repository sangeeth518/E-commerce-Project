package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password"`
}

func (a *Admin) HashPassword(password string) (string, error) {

	byte, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	a.Password = string(byte)
	return a.Password, nil
}
