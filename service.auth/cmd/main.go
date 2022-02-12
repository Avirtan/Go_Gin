package main

import (
	"avirtan/work_craft/config"
	docs "avirtan/work_craft/service.auth/api/v1"
	AuthRoute "avirtan/work_craft/service.auth/routes"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var addres string = "0.0.0.0"
var port string = "3000"
var mode string

//инициализация env файла и переменных
func init() {
	err := godotenv.Load(".env")
	mode = os.Getenv("MODE")
	if err != nil {
		panic("NOT LOAD ENV")
	}
}

// @title           Auth Work_Craft
// @version         1.0
// @description     This is a auth service.
// @contact.url   	http://work-craft.ru
// @host      		192.168.1.185
func main() {
	gin.SetMode(mode)
	app := gin.Default()
	app.SetTrustedProxies([]string{"0.0.0.0", "80.78.244.199", "localhost", "192.168.1.185"})
	app.Use(cors.Default())
	config.ConnectDatabasePostgres()
	config.ConnectRedis()
	routeApi := app.Group("api/v1/auth")
	AuthRoute.Routes(app, routeApi)
	if mode == "debug" {
		docs.SwaggerInfo.BasePath = "/api/v1/auth"
		routeApi.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	//fmt.Println("run server")
	srv := &http.Server{
		Addr:           addres + ":" + port,
		Handler:        app,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
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

//app.RunTLS(":3000", "/etc/nginx/ssl/domain.crt", "/etc/nginx/ssl/domain.key")
//http.ListenAndServeTLS(":3000", "/etc/nginx/ssl/domain.crt", "/etc/nginx/ssl/domain.key", app)
//app.Run(addres + ":" + port)

func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		//start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		//duration := util.GetDurationInMillseconds(start)
		entry := log.WithFields(log.Fields{
			//"client_ip":  util.GetClientIP(c),
			//"duration": duration,
			"method": c.Request.Method,
			"path":   c.Request.RequestURI,
			"status": c.Writer.Status(),
			//"user_id":    util.GetUserID(c),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
			// "api_version": util.ApiVersion,
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}
