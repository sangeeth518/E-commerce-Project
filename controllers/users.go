package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sangeeth518/E-commerce-Project/auth"
	"github.com/sangeeth518/E-commerce-Project/config"
	"github.com/sangeeth518/E-commerce-Project/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserSignup

func UserSignup(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		c.Abort()
		return
	}
	validationErr := validator.New().Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		return
	}
	config.DB.First(&user, "email=?", user.Email)
	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exist"})
		return
	}
	bytes, err := user.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to hash"})
		c.Abort()
		return
	}

	users := models.User{FirstName: user.FirstName, Lastname: user.Lastname, Email: user.Email, Password: bytes, PhoneNumber: user.PhoneNumber}
	result := config.DB.Create(&users)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create user"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sucess": "OK",
	})
}

// UserLogin

func UserLogin(c *gin.Context) {
	var user User

	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't bind"})
		return

	}
	var users models.User
	config.DB.First(&users, "email=?", user.Email)
	if users.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid email adress"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password dosen't match"})
		return
	}
	tokensting, err := auth.GenrateJWt(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	token := tokensting["access_token"]
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("userauthorization", token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusAccepted, gin.H{"Status": "true", "tokenstring": tokensting})

}
