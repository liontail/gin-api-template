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

	r.GET("/", mc.Hello)

	return r
}
