package AuthRoute

import (
	AuthController "avirtan/work_craft/service.auth/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine, route *gin.RouterGroup) {
	route.POST("/reg", AuthController.Registration)
	route.POST("/login", AuthController.Login)
}
