package ChatController

import (
	"avirtan/work_craft/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Chat godoc
// @Summary 	Chat
// @Description get all chat
// @Tags 		Chat
// @Success		200 {array} string
// @Failure		400 {integer} http.StatusBadRequest
// @Security 	ApiKeyAuth
// @Router 		/all [get]
func GetAllChat(ctx *gin.Context) {
	chats, err := config.DB_Redis.SMembers(ChatsKey).Result()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"err:": "error gets chats"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"chats": chats})
}
