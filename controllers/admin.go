package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/E-commerce-Project/auth"
	i "github.com/sangeeth518/E-commerce-Project/config"
	"github.com/sangeeth518/E-commerce-Project/models"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AdminSignup(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBind(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	i.DB.First(&admin, "email=?", admin.Email)
	if admin.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "This email already registered..",
		})
		return
	}
	bytes, err := admin.HashPassword(admin.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "failed to hash",
		})
		return
	}

	admins := models.Admin{Email: admin.Email, Password: bytes}
	record := i.DB.Create(&admins)
	if record.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create Admin",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func AdminLogin(c *gin.Context) {
	var admin Admin
	if err := c.ShouldBind(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}
	var admins models.Admin
	i.DB.First(&admins, "email=?", admin.Email)
	if admins.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid Email adress",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(admins.Password), []byte(admin.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "password doesn't match",
		})
		return
	}

	tokensring, err := auth.GenrateJWt(admin.Email)
	token := tokensring["access_token"]

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      "true",
		"tokenstring": tokensring,
	})

}

func AdminHome(c *gin.Context) {

	c.JSON(http.StatusAccepted, gin.H{
		"status": "Welcome to admin home page ",
	})
}

func AdminSignout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{
		"Message": "Admin Successfully Signed Out",
	})
}
