package AuthController

import (
	"net/http"

	jwtHandler "avirtan/work_craft/pkg/jwt"
	AuthModels "avirtan/work_craft/service.auth/models"

	"github.com/gin-gonic/gin"
)

// поля для регистрации
type RegistrationType struct {
	Login    string `json:"login" binding:"required" example:"root"`
	Password string `json:"password" binding:"required" example:"root1234"`
	Email    string `json:"email" binding:"required" example:"root@mail.ru"`
}

// Registration godoc
// @Summary 	registration user
// @Description registration user by json
// @Tags 		user
// @Accept 		json
// @Produce 	json
// @Param 		user body RegistrationType true "registration user"
// @Success		200 {string} string
// @Failure		400 {integer} http.StatusBadRequest
// @Router 		/reg [post]
func Registration(ctx *gin.Context) {
	var json RegistrationType
	user := &AuthModels.User{}
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	err = user.CreateUser(json.Login, json.Password, json.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := jwtHandler.CreateToken(jwtHandler.UserDate{Id: user.ID.String(), Login: user.Login})
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
