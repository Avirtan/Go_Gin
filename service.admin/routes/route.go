package AdminRoute

import (
	ControllerAdminProfile "avirtan/work_craft/service.admin/controllers/profile"

	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine, route *gin.RouterGroup) {
	profile := route.Group("/profile")
	{
		profile.GET("/users", ControllerAdminProfile.GetUsers)
		profile.GET("/user/:id", ControllerAdminProfile.GetUser)
		profile.PATCH("/user", ControllerAdminProfile.UpdateUser)
		profile.DELETE("/role", ControllerAdminProfile.DeleteRole)
		profile.POST("/role", ControllerAdminProfile.CreateRole)
		//profile.POST("/setRole", ControllerAdminProfile.SetRoleUser)
		profile.GET("/roles", ControllerAdminProfile.GetRoles)
		//profile.POST("/upload/:name", ControllerAdminProfile.UploadThings)
		profile.POST("/upload/:name", ControllerAdminProfile.UploadItem)
		profile.GET("/items", ControllerAdminProfile.GetAllItems)
		profile.DELETE("/items", ControllerAdminProfile.DeleteItem)
		profile.GET("/characters", ControllerAdminProfile.GetCharacters)
	}
}
