package web

import (
	// "fmt"
	"io"
	"os"

	"gin-demo/pkg/config"
	"gin-demo/pkg/web/controller"
	"gin-demo/pkg/web/middleware"

	"github.com/gin-gonic/gin"
)

// Router maps URI to handle function
var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.Use(middleware.GetLogger())
	// Router.Use(gin.Recovery())
	Router.Use(middleware.Recover())
}

// MapURI maps URI to funcs
func MapURI(conf *config.Config) {
	v := Router.Group(*conf.Server.Path)
	v.GET("/hello", controller.Hello)
	v.GET("/monitor", controller.Monitor)

	v.GET("/users/:token", controller.GetUserByID)
	v.PUT("/users/:token", controller.UpdateUser)
	v.GET("/users", controller.ListUsers)
	v.POST("/users", controller.AddUser)
}

func ConfigLog(conf *config.Config) {
	// logfileName := fmt.Sprintf("%s/gin-demo.log", *conf.Server.LogPath)
	logfile, _ := os.Create("gin-demo.log")
	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)
}
