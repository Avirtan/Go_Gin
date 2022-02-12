package ChatController

import (
	"avirtan/work_craft/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ключи для redis
const (
	ChatKey  = "chat:%s"
	ChatsKey = "chats"
)

//поля для создания чата
type Chat struct {
	Name string `json:"name" binding:"required" example:"Chan1"`
}

// Chat godoc
// @Summary 	Chat
// @Description creat chat by json
// @Tags 		Chat
// @Accept 		json
// @Produce 	json
// @Param 		chat body Chat true "chat"
// @Success		200 {string} string
// @Failure		400 {integer} http.StatusBadRequest
// @Security 	ApiKeyAuth
// @Router 		/create [post]
func CreateChat(ctx *gin.Context) {
	var json Chat
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if config.DB_Redis.SIsMember(ChatsKey, json.Name).Val() {
		ctx.JSON(http.StatusBadRequest, gin.H{"err:": "channels is exist"})
		return
	}
	if err := config.DB_Redis.SAdd(ChatsKey, json.Name).Err(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "chat create"})
}
