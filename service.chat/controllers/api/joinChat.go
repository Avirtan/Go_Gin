package ChatController

import (
	"avirtan/work_craft/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Chat godoc
// @Summary 	Chat
// @Description join to chat by json
// @Tags 		Chat
// @Accept 		json
// @Produce 	json
// @Param 		chat body Chat true "chat"
// @Success		200 {string} string
// @Failure		400 {integer} http.StatusBadRequest
// @Security 	ApiKeyAuth
// @Router 		/join [post]
func JoinChat(ctx *gin.Context) {
	var json Chat
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !config.DB_Redis.SIsMember(ChatsKey, json.Name).Val() {
		ctx.JSON(http.StatusBadRequest, gin.H{"err:": "channels not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": fmt.Sprintf("join to %v", json.Name)})
}
