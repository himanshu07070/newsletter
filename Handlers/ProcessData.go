package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	input "github.com/himanshu07070/newsletter/Input"
	utils "github.com/himanshu07070/newsletter/Utils"
)

func Newsletter(c *gin.Context) {

	var metaData input.MetaData
	err := c.BindJSON(&metaData)
	if err != nil {
		utils.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, "invalid input")
		return
	}
	dataChannel := c.MustGet("channel").(chan input.MetaData)
	dataChannel <- metaData

	utils.Logger.Info("data sent to channel")
	c.JSON(http.StatusOK, "sucess")
}
