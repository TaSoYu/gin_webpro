package middlewares

import (
	"exchangeapp/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization header",
			})
			ctx.Abort()
			return
		}
		username, err := utils.ParseJWT(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			ctx.Abort()
			return
		}
		//调试
		log.Printf("用户: " + username + " 通过中间件检测")

		ctx.Set("username", username)
		ctx.Next()
	}
}
