package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/E-commerce-Project/config"
	"github.com/sangeeth518/E-commerce-Project/initializers"
	"github.com/sangeeth518/E-commerce-Project/routes"
)

var app = gin.Default()

func init() {

	initializers.LoadEnv()
	config.DbConnect()

}
func main() {
	port := os.Getenv("port")
	routes.AdminRoutes(app)
	routes.UserRoutes(app)
	app.Run(port)

}
