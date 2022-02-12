package AuthRoute

import (
	ProfileController "avirtan/work_craft/service.profile/controllers"
	Middleware "avirtan/work_craft/service.profile/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine, route *gin.RouterGroup) {
	profile := route.Group("", Middleware.AuthorizeJWT())
	{
		profile.GET("/", ProfileController.GetProfile)
		profile.POST("/character", ProfileController.CreateCharacter)
		profile.GET("/items/user", ProfileController.GetItemsUser)
		profile.GET("/items", ProfileController.GetAllItems)
	}
}
