package main

import (
	"avirtan/work_craft/config"
	docs "avirtan/work_craft/service.profile/api/v1"
	Middleware "avirtan/work_craft/service.profile/middleware"
	ProfileRoute "avirtan/work_craft/service.profile/routes"
	"fmt"

	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var addres string = "0.0.0.0"
var port string = "3001"
var mode string

//инициализация env файла и переменных
func init() {
	err := godotenv.Load(".env")
	mode = os.Getenv("MODE")
	if err != nil {
		panic("NOT LOAD ENV")
	}
}

// @title           			Profile Work_Craft
// @version         			1.0
// @description     			This is a profile service.
// @contact.url   				http://work-craft.ru
// @host      					192.168.1.185
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	gin.SetMode(mode)
	app := gin.Default()
	app.SetTrustedProxies([]string{"0.0.0.0", "80.78.244.199", "localhost", "192.168.1.185"})
	app.Use(Middleware.CORS())
	config.ConnectDatabasePostgres()
	config.ConnectRedis()
	routeApi := app.Group("api/v1/profile")
	ProfileRoute.Routes(app, routeApi)
	if mode == "debug" {
		docs.SwaggerInfo.BasePath = "/api/v1/profile"
		routeApi.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	} else {
		app.Use(Middleware.LogMiddleware())
	}
	srv := &http.Server{
		Addr:           addres + ":" + port,
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := srv.ListenAndServeTLS("/etc/nginx/ssl/domain.crt", "/etc/nginx/ssl/domain.key"); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()
	fmt.Println("server run " + addres + ":" + port)
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
