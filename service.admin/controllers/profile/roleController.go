package ControllerAdminProfile

import (
	profile_models "avirtan/work_craft/service.profile/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleAdd struct {
	Name string `json:"name" example:"test_role"`
}

type RoleDelete struct {
	ID uint `json:"id"`
}

type RoleSet struct {
	IdUser string `json:"idUser"`
	IdRole uint   `json:"idRole"`
}

func GetRoles(ctx *gin.Context) {
	roles, err := profile_models.GetAllRole()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"roles": roles})
}

func DeleteRole(ctx *gin.Context) {
	json := RoleDelete{}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	role := &profile_models.Role{ID: json.ID}
	fmt.Println(role)
	err := role.DeleteRole()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "delete"})
}

func CreateRole(ctx *gin.Context) {
	json := RoleAdd{}
	//fmt.Println(ctx.Request.Body)
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	role := &profile_models.Role{Name: json.Name}
	err := role.AddRole()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"role": role})
}

