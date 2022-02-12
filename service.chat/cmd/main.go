package main

import (
	"avirtan/work_craft/config"
	docs "avirtan/work_craft/service.chat/api/v1"
	ChatRoute "avirtan/work_craft/service.chat/routes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var addres string = "0.0.0.0"
var port string = "3002"
var mode string

func init() {
	err := godotenv.Load(".env")
	mode = os.Getenv("MODE")
	if err != nil {
		panic("NOT LOAD ENV")
	}
}

// @title          				Chat Work_Craft
// @version        				1.0
// @description    				This is a chat service.
// @contact.url    				http://work-craft.ru
// @host      					0.0.0.0:3001
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	gin.SetMode(mode)
	app := gin.Default()
	app.SetTrustedProxies([]string{"0.0.0.0", "80.78.244.199", "localhost"})
	app.Use(cors.Default())
	config.ConnectDatabasePostgres()
	config.ConnectRedis()
	routeApi := app.Group("api/v1/chat")
	ChatRoute.Routes(app, routeApi)
	if mode == "debug" {
		docs.SwaggerInfo.BasePath = "/api/v1/chat"
		routeApi.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	srv := &http.Server{
		Addr:           addres + ":" + port,
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("server run " + addres + ":" + port)
	go func() {
		if err := srv.ListenAndServeTLS("/etc/nginx/ssl/domain.crt", "/etc/nginx/ssl/domain.key"); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
