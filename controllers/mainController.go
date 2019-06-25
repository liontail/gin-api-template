package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liontail/go-api-template/models"
)

type MainController struct {
}

func (con *MainController) Hello(c *gin.Context) {
	c.JSON(200, models.Object{
		Message: "Hello World",
	})
	return
}
