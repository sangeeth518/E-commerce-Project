package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

// <<<<<<<<<<<<<<<<<<<<<<<<<< Admin Signup >>>>>>>>>>>>>>>>>>>>>>>>>>>>>

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

// <<<<<<<<<<<<<<<<<<<<<<Admin Login>>>>>>>>>>>>>>>>>>>>>>>>>

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

// Get users details for authenticated admins

func GetUsers(c *gin.Context) {
	pagestr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "page not given in right format"})
	}
	pagesize, err := strconv.Atoi(c.DefaultQuery("count", "3"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user count": "not in right format"})
	}
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * pagesize
	var userdetails models.UserDetailsAtAdmin
	if err := i.DB.Raw("select id , email, phone_number,block_status from users order by id asc limit ? offset ?", pagesize, offset).Scan(&userdetails).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "error"})
	}
	c.JSON(200, gin.H{"user": userdetails})

}

func GetUserbyId(id string) (models.User, error) {
	user_id, err := strconv.Atoi(id)
	if err != nil {
		return models.User{}, err
	}
	var count int
	if err := i.DB.Raw("select count(*) from users where id = ?", user_id).Scan(&count).Error; err != nil {
		return models.User{}, err
	}
	if count < 1 {
		return models.User{}, errors.New("user for the given id does not exist")
	}
	query := fmt.Sprintf("select * from users where id = '%d'", user_id)
	var userdetails models.User
	if err := i.DB.Raw(query).Scan(&userdetails).Error; err != nil {
		return models.User{}, err
	}
	return userdetails, nil

}
func BlockUser(c *gin.Context) {

	params := c.Param("id")
	user, err := GetUserbyId(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user for the given id": "does not exist"})
		return

	}
	if user.Block_status {
		c.JSON(http.StatusBadRequest, gin.H{"already ": "blocked"})
		return
	} else {
		user.Block_status = true
	}
	err = UpdateBlockUserByID(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"couldnt": "block"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user ": "blocked"})

}
func UnblockUser(c *gin.Context) {
	params := c.Param("id")
	user, err := GetUserbyId(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "given id does not exist"})
		return
	}
	if user.Block_status {
		user.Block_status = false
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"already ": "unblocked"})
		return
	}

	err = UpdateBlockUserByID(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"couldnt": "unblock"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "unblocked sucessfully"})

}

// function which will both block and unblock a user
func UpdateBlockUserByID(user models.User) error {
	err := i.DB.Exec("update users set block_status = ? where id = ?", user.Block_status, user.ID).Error
	if err != nil {
		return err
	}
	return nil

}
