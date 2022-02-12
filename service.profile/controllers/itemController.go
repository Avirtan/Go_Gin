package ProfileController

import (
	profile_models "avirtan/work_craft/service.profile/models"
	profileCharacter "avirtan/work_craft/service.profile/models/characterItems"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// Profile godoc
// @Summary 	GetItems
// @Description GetItmes
// @Tags 		items
// @Success		200  {string} string
// @Failure		400  {integer} http.StatusBadRequest
// @Security 	ApiKeyAuth
// @Router 		/items [get]
func GetAllItems(ctx *gin.Context) {
	eyes, err := profileCharacter.GetAllEyes()
	if HandlerError(ctx, err) {
		return
	}
	hair, err := profileCharacter.GetAllHair()
	if HandlerError(ctx, err) {
		return
	}
	mouth, err := profileCharacter.GetAllMouth()
	if HandlerError(ctx, err) {
		return
	}
	nose, err := profileCharacter.GetAllNose()
	if HandlerError(ctx, err) {
		return
	}
	other, err := profileCharacter.GetAllOther()
	if HandlerError(ctx, err) {
		return
	}
	bottom, err := profileCharacter.GetAllBottom()
	if HandlerError(ctx, err) {
		return
	}
	shoes, err := profileCharacter.GetAllShoes()
	if HandlerError(ctx, err) {
		return
	}
	top, err := profileCharacter.GetAllTop()
	if HandlerError(ctx, err) {
		return
	}
	weapon, err := profileCharacter.GetAllWeapon()
	if HandlerError(ctx, err) {
		return
	}
	items := make(map[string]interface{})
	items["eyes"] = eyes
	items["hair"] = hair
	items["mouth"] = mouth
	items["nose"] = nose
	items["other"] = other
	items["bottom"] = bottom
	items["shoes"] = shoes
	items["top"] = top
	items["weapons"] = weapon
	ctx.JSON(http.StatusOK, gin.H{"items": items})
}

// Profile godoc
// @Summary 	GetUserItems
// @Description GetUserItmes
// @Tags 		items
// @Success		200  {string} string
// @Failure		400  {integer} http.StatusBadRequest
// @Security 	ApiKeyAuth
// @Router 		/items/user [get]
func GetItemsUser(ctx *gin.Context) {
	id, _ := ctx.Get("id")
	userId, err := uuid.FromString(fmt.Sprintf("%v", id))
	if HandlerError(ctx, err) {
		return
	}
	user, err := profile_models.GetUserById(userId.String())
	if HandlerError(ctx, err) {
		return
	}
	err = user.GetCharacter()
	if HandlerError(ctx, err) {
		return
	}
	eyes, err := profileCharacter.GetEyesById(user.Character.EyeID)
	if HandlerError(ctx, err) {
		return
	}
	hair, err := profileCharacter.GetHairById(user.Character.HairID)
	if HandlerError(ctx, err) {
		return
	}
	mouth, err := profileCharacter.GetMouthById(user.Character.MouthID)
	if HandlerError(ctx, err) {
		return
	}
	nose, err := profileCharacter.GetNoseById(user.Character.NoseID)
	if HandlerError(ctx, err) {
		return
	}
	other, err := profileCharacter.GetOtherById(user.Character.OtherID)
	if HandlerError(ctx, err) {
		return
	}
	bottom, err := profileCharacter.GetBottomById(user.Character.BottomID)
	if HandlerError(ctx, err) {
		return
	}
	shoes, err := profileCharacter.GetShoeById(user.Character.ShoeID)
	if HandlerError(ctx, err) {
		return
	}
	top, err := profileCharacter.GetTopById(user.Character.TopID)
	if HandlerError(ctx, err) {
		return
	}
	weapon, err := profileCharacter.GetWeaponById(user.Character.WeaponID)
	if HandlerError(ctx, err) {
		return
	}
	items := make(map[string]interface{})
	items["eyes"] = eyes
	items["hair"] = hair
	items["mouth"] = mouth
	items["nose"] = nose
	items["other"] = other
	items["bottom"] = bottom
	items["shoes"] = shoes
	items["top"] = top
	items["weapons"] = weapon
	ctx.JSON(http.StatusOK, gin.H{"items": items})
}
