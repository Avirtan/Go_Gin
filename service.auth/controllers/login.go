package AuthController

import (
	jwtHandler "avirtan/work_craft/pkg/jwt"
	AuthModels "avirtan/work_craft/service.auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// поля для авторизации
type LoginType struct {
	Login    string `json:"login" example:"root"`
	Password string `json:"password" binding:"required,min=5" example:"root1234"`
}

// Login godoc
// @Summary 	login user
// @Description login user by json
// @Tags 		user
// @Accept 		json
// @Produce 	json
// @Param 		user body LoginType true "login user"
// @Success		200  {string} string
// @Failure		400  {integer} http.StatusBadRequest
// @Router 		/login [post]
func Login(ctx *gin.Context) {
	var json LoginType
	user := &AuthModels.User{}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := user.FindUser(json.Login, json.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := jwtHandler.CreateToken(jwtHandler.UserDate{Login: user.Login, Id: user.ID.String()})
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
