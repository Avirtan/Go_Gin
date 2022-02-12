package ProfileController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerError(ctx *gin.Context, err error) bool {
	if err != nil {
		//ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return true
	}
	return false
}
