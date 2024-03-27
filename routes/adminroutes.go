package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/E-commerce-Project/controllers"
	"github.com/sangeeth518/E-commerce-Project/middleware"
)

func AdminRoutes(c *gin.Engine) {
	admin := c.Group("/admin")
	{
		admin.POST("/signup", controllers.AdminSignup)
		admin.POST("/login", controllers.AdminLogin)
		admin.GET("/home", middleware.Adminauth(), controllers.AdminHome)
		admin.GET("/signout", controllers.AdminSignout, middleware.Adminauth())
	}
}
