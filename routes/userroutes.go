package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/E-commerce-Project/controllers"
)

func UserRoutes(c *gin.Engine) {
	user := c.Group("/user")
	{
		user.POST("/signup", controllers.UserSignup)
		user.POST("/login", controllers.UserLogin)
		user.POST("/addadress/", controllers.AddAddress)
	}

}
