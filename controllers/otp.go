package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/E-commerce-Project/helper"
	"github.com/sangeeth518/E-commerce-Project/models"
)

func SndOtp(c *gin.Context) {
	var phone models.OTPData
	if err := c.BindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "couldn't bind"})
		return
	}
	helper.TwilioSetup(os.Getenv("ACCOUNT_SID"), os.Getenv("AUTH_TOKEN"))
	_, err := helper.TwilioSndOtp(phone.PhoneNumber, os.Getenv("SERVICE_SID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "while generating otp"})
		return

	}
	c.JSON(http.StatusOK, "otp send succesfullly")
}
