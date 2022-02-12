package middlewareJwt

import (
	jwtService "avirtan/work_craft/pkg/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// метод на предварительную аутентификацию пользователя по токену
func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) == 0 {
			//fmt.Println("notFound")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "not found token",
			})
			return
		}
		if _, ok := jwtService.ValidateToken(authHeader); !ok {
			fmt.Println("error token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token not valid",
			})
			return
		}
		ctx.Next()
	}
}
