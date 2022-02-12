package ProfileController

import (
	profile_models "avirtan/work_craft/service.profile/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProfile(ctx *gin.Context) {

}

// Profile godoc
// @Summary 	GetProfile
// @Description GetProfile
// @Tags 		profile
// @Success		200  {object} profile_models.User
// @Failure		400  {integer} http.StatusBadRequest
// @Security 	ApiKeyAuth
// @Router 		/ [get]
func GetProfile(ctx *gin.Context) {
	id, _ := ctx.Get("id")
	user, err := profile_models.GetUserById(fmt.Sprintf("%v", id))
	if HandlerError(ctx, err) {
		return
	}
	err = user.GetStatistic()
	if HandlerError(ctx, err) {
		return
	}
	err = user.GetCharacter()
	if HandlerError(ctx, err) {
		return
	}
	if user.Character.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{"error": "not found character"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": user})
}
