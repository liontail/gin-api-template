package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liontail/go-api-template/controllers"
)

func SetupRouter() *gin.Engine {
	runmode := "debug"

	gin.SetMode(runmode)

	r := gin.Default()

	mc := controllers.MainController{}

	r.GET("/get", mc.Get)
	r.POST("/post", mc.Post)
	return r
}
