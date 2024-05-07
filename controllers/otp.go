package controllers

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/E-commerce-Project/config"
	"github.com/sangeeth518/E-commerce-Project/helper"
	"github.com/sangeeth518/E-commerce-Project/models"
	"gorm.io/gorm"
)

func SndOtp(c *gin.Context) {
	var phone models.OTPData
	if err := c.BindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, "couldn't bind")
		return
	}
	ok := FindUserByPhone(phone.Number)
	if ok != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "user does not exist"})
		return
	}

	helper.TwilioSetup(os.Getenv("ACCOUNT_SID"), os.Getenv("AUTH_TOKEN"))
	_, err := helper.TwilioSndOtp(phone.Number, os.Getenv("SERVICE_SID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERR": "COULDN'T SND OTP"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"otp": "snd succesfully"})
}

func FindUserByPhone(phone string) error {
	var user models.User
	result := config.DB.Where(&models.User{PhoneNumber: phone}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("err")
		}
		return result.Error
	}
	return nil
}
