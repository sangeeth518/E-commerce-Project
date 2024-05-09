package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/E-commerce-Project/controllers"
	"github.com/sangeeth518/E-commerce-Project/middleware"
)

func UserRoutes(c *gin.Engine) {
	user := c.Group("/user")
	{
		user.POST("/signup", controllers.UserSignup)
		user.POST("/login", controllers.UserLogin)
		user.POST("/otplogin", controllers.SndOtp)
		user.POST("/verifyotp", controllers.VerifyOtp)
		user.POST("/forgotpass", controllers.ForgotPasswordSend)
		user.POST("/addadress/", middleware.UserAuth(), controllers.AddAddress)
		user.GET("/getaddress", middleware.UserAuth(), controllers.GetAdresses)
		user.GET("/userdetails", middleware.UserAuth(), controllers.GetUserDetails)
		user.PUT("/change-password", middleware.UserAuth(), controllers.ChangePassword)

	}

}
