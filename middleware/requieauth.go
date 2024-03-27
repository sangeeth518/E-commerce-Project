package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/E-commerce-Project/auth"
)

func Adminauth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenstring, err := ctx.Cookie("Authorization")
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		if tokenstring == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}
		err = auth.Validtoken(tokenstring)
		if err != nil {
			ctx.AbortWithStatus(401)
		}

		ctx.Next()

	}
}
