package ControllerAdminProfile

import (
	Handler "avirtan/work_craft/service.admin/controllers"
	profile_models "avirtan/work_craft/service.profile/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCharacters(ctx *gin.Context) {
	characters, err := profile_models.GetAllCharacter()
	if Handler.HandlerError(ctx, err) {
		return
	}
	users, err := profile_models.GetAllUser()
	if Handler.HandlerError(ctx, err) {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"characters": characters, "users": users})
}
