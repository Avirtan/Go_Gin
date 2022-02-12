package ControllerAdminProfile

import (
	Handler "avirtan/work_craft/service.admin/controllers"
	profileCharacter "avirtan/work_craft/service.profile/models/characterItems"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// func UploadThings(ctx *gin.Context) {
// 	name := ctx.Param("name")
// 	fileTypes := []string{"image/jpeg", "image/png"}
// 	file, err := ctx.FormFile("img")
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 		return
// 	}
// 	fileType := file.Header["Content-Type"][0]
// 	found := false
// 	for _, val := range fileTypes {
// 		if val == fileType {
// 			found = true
// 			break
// 		}
// 	}
// 	if !found || file.Size > (8<<20) {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("error type file only  %v", fileTypes)})
// 		return
// 	}
// 	path := os.Getenv("PATH_TO_SAVE_FILE")
// 	fullPath := fmt.Sprintf("%s%s/%s", path, name, filepath.Base(file.Filename))
// 	_, err = os.Stat(fullPath)
// 	if err != nil {
// 		fmt.Println("create")
// 		os.MkdirAll(fmt.Sprintf("%s%s/", path, name), os.ModePerm)
// 	}
// 	switch name {
// 	case "eye":
// 		fmt.Println(file.Filename)
// 		eye := &profileCharacter.Eye{
// 			Name: file.Filename,
// 			Path: "img/" + name + "/" + file.Filename,
// 		}
// 		err = eye.Create()
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 			return
// 		}
// 	case "hair":
// 		hair := &profileCharacter.Hair{
// 			Name: file.Filename,
// 			Path: "img/" + name + "/" + file.Filename,
// 		}
// 		err = hair.Create()
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 			return
// 		}
// 	case "mouth":
// 		mouth := &profileCharacter.Mouth{
// 			Name: file.Filename,
// 			Path: "img/" + name + "/" + file.Filename,
// 		}
// 		err = mouth.Create()
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 			return
// 		}
// 	case "nose":
// 		nose := &profileCharacter.Nose{
// 			Name: file.Filename,
// 			Path: "img/" + name + "/" + file.Filename,
// 		}
// 		err = nose.Create()
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 			return
// 		}
// 	case "other":
// 		other := &profileCharacter.Other{
// 			Name: file.Filename,
// 			Path: "img/" + name + "/" + file.Filename,
// 		}
// 		err = other.Create()
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 			return
// 		}
// 	case "leg":
// 		leg := &profileCharacter.Leg{
// 			Name: file.Filename,
// 			Path: "img/" + name + "/" + file.Filename,
// 		}
// 		err = leg.Create()
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 			return
// 		}
// 	case "shoe":
// 		shoe := &profileCharacter.Shoe{
// 			Name: file.Filename,
// 			Path: "img/" + name + "/" + file.Filename,
// 		}
// 		err = shoe.Create()
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 			return
// 		}
// 	case "top":
// 		top := &profileCharacter.top{
// 			Name: file.Filename,
// 			Path: "img/" + name + "/" + file.Filename,
// 		}
// 		err = top.Create()
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 			return
// 		}
// 	case "weapon":
// 		weapon := &profileCharacter.Weapon{
// 			Name: file.Filename,
// 			Path: "img/" + name + "/" + file.Filename,
// 		}
// 		err = weapon.Create()
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 			return
// 		}
// 	}
// 	if err := ctx.SaveUploadedFile(file, fullPath); err != nil {
// 		ctx.String(http.StatusBadRequest, "upload file err: %s", err.Error())
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
// }

func UploadItem(ctx *gin.Context) {
	name := ctx.Param("name")
	fileTypes := []string{"image/jpeg", "image/png"}
	form, err := ctx.MultipartForm()
	if Handler.HandlerError(ctx, err) {
		return
	}
	files, ok := form.File["imgs[]"]
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "not found file"})
		return
	}
	for _, file := range files {
		//fmt.Println(file.Filename)
		fileType := file.Header["Content-Type"][0]
		found := false
		for _, val := range fileTypes {
			if val == fileType {
				found = true
				break
			}
		}
		if !found || file.Size > (8<<20) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("error type file only  %v", fileTypes)})
			return
		}
		path := os.Getenv("PATH_TO_SAVE_FILE")
		fullPath := fmt.Sprintf("%s%s/%s", path, name, filepath.Base(file.Filename))
		_, err = os.Stat(fullPath)
		if err != nil {
			fmt.Println("create")
			os.MkdirAll(fmt.Sprintf("%s%s/", path, name), os.ModePerm)
		}
		switch name {
		case "eye":
			eye := &profileCharacter.Eye{
				Name: file.Filename,
				Path: "img/" + name + "/" + file.Filename,
			}
			err = eye.Create()
			if Handler.HandlerError(ctx, err) {
				return
			}
		case "hair":
			hair := &profileCharacter.Hair{
				Name: file.Filename,
				Path: "img/" + name + "/" + file.Filename,
			}
			err = hair.Create()
			if Handler.HandlerError(ctx, err) {
				return
			}
		case "mouth":
			mouth := &profileCharacter.Mouth{
				Name: file.Filename,
				Path: "img/" + name + "/" + file.Filename,
			}
			err = mouth.Create()
			if Handler.HandlerError(ctx, err) {
				return
			}
		case "nose":
			nose := &profileCharacter.Nose{
				Name: file.Filename,
				Path: "img/" + name + "/" + file.Filename,
			}
			err = nose.Create()
			if Handler.HandlerError(ctx, err) {
				return
			}
		case "other":
			other := &profileCharacter.Other{
				Name: file.Filename,
				Path: "img/" + name + "/" + file.Filename,
			}
			err = other.Create()
			if Handler.HandlerError(ctx, err) {
				return
			}
		case "bottom":
			bottom := &profileCharacter.Bottom{
				Name: file.Filename,
				Path: "img/" + name + "/" + file.Filename,
			}
			err = bottom.Create()
			if Handler.HandlerError(ctx, err) {
				return
			}
		case "shoe":
			shoe := &profileCharacter.Shoe{
				Name: file.Filename,
				Path: "img/" + name + "/" + file.Filename,
			}
			err = shoe.Create()
			if Handler.HandlerError(ctx, err) {
				return
			}
		case "top":
			top := &profileCharacter.Top{
				Name: file.Filename,
				Path: "img/" + name + "/" + file.Filename,
			}
			err = top.Create()
			if Handler.HandlerError(ctx, err) {
				return
			}
		case "weapon":
			weapon := &profileCharacter.Weapon{
				Name: file.Filename,
				Path: "img/" + name + "/" + file.Filename,
			}
			err = weapon.Create()
			if Handler.HandlerError(ctx, err) {
				return
			}
		}
		err := ctx.SaveUploadedFile(file, fullPath)
		if Handler.HandlerError(ctx, err) {
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
