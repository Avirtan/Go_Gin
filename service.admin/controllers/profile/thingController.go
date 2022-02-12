package ControllerAdminProfile

import (
	Handler "avirtan/work_craft/service.admin/controllers"
	profileCharacter "avirtan/work_craft/service.profile/models/characterItems"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ThingType struct {
	Id   uint   `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

func GetAllItems(ctx *gin.Context) {
	eyes, err := profileCharacter.GetAllEyes()
	if Handler.HandlerError(ctx, err) {
		return
	}
	hair, err := profileCharacter.GetAllHair()
	if Handler.HandlerError(ctx, err) {
		return
	}
	mouth, err := profileCharacter.GetAllMouth()
	if Handler.HandlerError(ctx, err) {
		return
	}
	nose, err := profileCharacter.GetAllNose()
	if Handler.HandlerError(ctx, err) {
		return
	}
	other, err := profileCharacter.GetAllOther()
	if Handler.HandlerError(ctx, err) {
		return
	}
	bottom, err := profileCharacter.GetAllBottom()
	if Handler.HandlerError(ctx, err) {
		return
	}
	shoes, err := profileCharacter.GetAllShoes()
	if Handler.HandlerError(ctx, err) {
		return
	}
	top, err := profileCharacter.GetAllTop()
	if Handler.HandlerError(ctx, err) {
		return
	}
	weapon, err := profileCharacter.GetAllWeapon()
	if Handler.HandlerError(ctx, err) {
		return
	}
	things := make(map[string]interface{})
	things["eyes"] = eyes
	things["hair"] = hair
	things["mouth"] = mouth
	things["nose"] = nose
	things["other"] = other
	things["bottom"] = bottom
	things["shoes"] = shoes
	things["top"] = top
	things["weapons"] = weapon
	ctx.JSON(http.StatusOK, gin.H{"items": things})
}

func DeleteItem(ctx *gin.Context) {
	json := &ThingType{}
	err := ctx.ShouldBindJSON(&json)
	if Handler.HandlerError(ctx, err) {
		return
	}
	path := os.Getenv("PATH_TO_SAVE_FILE")
	switch json.Type {
	case "eye":
		eye := &profileCharacter.Eye{ID: json.Id}
		err := eye.Delete()
		if Handler.HandlerError(ctx, err) {
			return
		}
	case "hair":
		hair := &profileCharacter.Hair{ID: json.Id}
		err := hair.Delete()
		if Handler.HandlerError(ctx, err) {
			return
		}
	case "mouth":
		mouth := &profileCharacter.Mouth{ID: json.Id}
		err := mouth.Delete()
		if Handler.HandlerError(ctx, err) {
			return
		}
	case "nose":
		nose := &profileCharacter.Nose{ID: json.Id}
		err := nose.Delete()
		if Handler.HandlerError(ctx, err) {
			return
		}
	case "other":
		other := &profileCharacter.Other{ID: json.Id}
		err := other.Delete()
		if Handler.HandlerError(ctx, err) {
			return
		}
	case "bottom":
		bottom := &profileCharacter.Bottom{ID: json.Id}
		err := bottom.Delete()
		if Handler.HandlerError(ctx, err) {
			return
		}
	case "shoe":
		shoe := &profileCharacter.Shoe{ID: json.Id}
		err := shoe.Delete()
		if Handler.HandlerError(ctx, err) {
			return
		}
	case "top":
		top := &profileCharacter.Top{ID: json.Id}
		err := top.Delete()
		if Handler.HandlerError(ctx, err) {
			return
		}
	case "weapon":
		weapon := &profileCharacter.Weapon{ID: json.Id}
		err := weapon.Delete()
		if Handler.HandlerError(ctx, err) {
			return
		}
	}
	err = os.Remove(path + json.Type + "/" + json.Name)
	if Handler.HandlerError(ctx, err) {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
