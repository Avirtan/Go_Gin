package ProfileController

import (
	profile_models "avirtan/work_craft/service.profile/models"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type CharacterType struct {
	Eyes   uint `json:"eyes"    binding:"required" example:"0"`
	Hair   uint `json:"hair"   binding:"required" example:"0"`
	Mouth  uint `json:"mouth"  binding:"required" example:"0"`
	Nose   uint `json:"nose"   binding:"required" example:"0"`
	Other  uint `json:"other" `
	Bottom uint `json:"bottom" binding:"required" example:"0"`
	Shoes  uint `json:"shoes"   binding:"required" example:"0"`
	Top    uint `json:"top"    binding:"required" example:"0"`
	Weapon uint `json:"weapon"`
}

// Character godoc
// @Summary 	CreateCharacter
// @Description CreateCharacter
// @Tags 		character
// @Param 		chat body CharacterType true "character"
// @Success		200  {object} profile_models.Character
// @Failure		400  {integer} http.StatusBadRequest
// @Security 	ApiKeyAuth
// @Router 		/character [post]
func CreateCharacter(ctx *gin.Context) {
	var json CharacterType
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
	if user.Character.ID != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "character found"})
		return
	}
	character := profile_models.Character{
		EyeID:    json.Eyes,
		HairID:   json.Hair,
		MouthID:  json.Mouth,
		NoseID:   json.Nose,
		OtherID:  json.Other,
		BottomID: json.Bottom,
		ShoeID:   json.Shoes,
		TopID:    json.Top,
		UserID:   userId,
	}
	err = character.Create()
	if HandlerError(ctx, err) {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": character})
}
