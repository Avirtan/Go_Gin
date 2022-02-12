package ControllerAdminProfile

import (
	Handler "avirtan/work_craft/service.admin/controllers"
	user_models "avirtan/work_craft/service.profile/models"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type UserSerialize struct {
	ID           uuid.UUID `json:"Id"`
	Login        string    `json:"Login"`
	Email        string    `json:"Email"`
	Verification bool      `json:"Verification"`
	CreatedAt    time.Time `json:"Create"`
	IconID       *uint     `json:"Icon"`
}

func Response(user user_models.User) UserSerialize {
	US := UserSerialize{
		ID:           user.ID,
		Login:        user.Login,
		Email:        user.Email,
		Verification: user.Verification,
		CreatedAt:    user.CreatedAt,
		IconID:       user.IconID,
	}
	return US
}

func GetUsers(ctx *gin.Context) {
	users, err := user_models.GetAllUser()
	if Handler.HandlerError(ctx, err) {
		return
	}
	usersSerialize := make([]UserSerialize, 0)
	for _, user := range users {
		usersSerialize = append(usersSerialize, Response(user))
	}
	ctx.JSON(http.StatusOK, gin.H{"users": usersSerialize})
}

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := user_models.GetUserById(id)
	user.GetRoles()
	if Handler.HandlerError(ctx, err) {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(ctx *gin.Context) {
	json := &user_models.User{}
	err := ctx.ShouldBindJSON(&json)
	if Handler.HandlerError(ctx, err) {
		return
	}
	user, err := user_models.GetUserById(json.ID.String())
	if Handler.HandlerError(ctx, err) {
		return
	}
	err = user.UpdateUser(json)
	if Handler.HandlerError(ctx, err) {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
