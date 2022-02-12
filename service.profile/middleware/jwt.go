package middlewareJwt

import (
	jwtService "avirtan/work_craft/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// метод на предварительную аутентификацию пользователя по токену
func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "not found token",
			})
			return
		}
		data, ok := jwtService.ValidateToken(authHeader)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token not valid",
			})
			return
		}
		//fmt.Println(data["id"])
		ctx.Set("id", data["id"])
		ctx.Next()
	}
}
