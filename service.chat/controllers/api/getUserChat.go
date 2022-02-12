package ChatController

import (
	"avirtan/work_craft/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Chat godoc
// @Summary 	Chat
// @Description get user from chat by json
// @Tags 		Chat
// @Accept 		json
// @Produce 	json
// @Param 		chat body Chat true "chat"
// @Success		200 {string} string
// @Failure		400 {integer} http.StatusBadRequest
// @Security 	ApiKeyAuth
// @Router 		/getUser [post]
func GetUsersChat(ctx *gin.Context) {
	var json Chat
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !config.DB_Redis.SIsMember(ChatsKey, json.Name).Val() {
		ctx.JSON(http.StatusBadRequest, gin.H{"err:": "chat not exist"})
		return
	}
	nameChat := fmt.Sprintf(ChatKey, json.Name)
	users, err := config.DB_Redis.SMembers(nameChat).Result()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err:": "error gets users"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
