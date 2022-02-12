package ChatRoute

import (
	middlewareJwt "avirtan/work_craft/middleware"
	ChatController "avirtan/work_craft/service.chat/controllers/api"
	ChatControllerWS "avirtan/work_craft/service.chat/controllers/ws"

	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine, route *gin.RouterGroup) {

	r := route.Group("", middlewareJwt.AuthorizeJWT())
	{
		r.POST("/create", ChatController.CreateChat)
		r.DELETE("/delete", ChatController.DeleteChat)
		r.GET("/all", ChatController.GetAllChat)
		r.POST("/getUser", ChatController.GetUsersChat)
		r.POST("/join", ChatController.JoinChat)
	}
	route.GET("/ws/chat/:name/:token", func(ctx *gin.Context) {
		ChatControllerWS.ChatWebSocketHandler(ctx.Writer, ctx.Request, ctx)
	})
}
