package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/liontail/go-api-template/models"
)

type MainController struct {
}

func (con *MainController) Get(c *gin.Context) {
	c.JSON(200, models.Object{
		Message: "Hello World",
	})
	return
}

func (con *MainController) Post(c *gin.Context) {
	var body interface{}
	if err := c.ShouldBind(&body); err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(200, body)
	return
}
