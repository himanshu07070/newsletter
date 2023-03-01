package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/himanshu07070/newsletter/Database"
	input "github.com/himanshu07070/newsletter/Input"
	utils "github.com/himanshu07070/newsletter/Utils"
)

func UnsubscribeEmails(c *gin.Context) {
	utils.Logger.Info("Inside UnsubscribeEmails...")
	ctx := c.Request.Context()
	var user input.UserEmail
	err := c.BindJSON(&user)
	if err != nil {
		utils.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, "invalid input")
		return
	}
	_, exist := database.CheckUserEmail(user.Email, ctx)
	if !exist {
		utils.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, "user does not exist")
		return
	}

	updateData := make(map[string]interface{})
	updateData["subscribe"] = false
	err = database.UnsubscribeEmailWithEmailId(user.Email, updateData, ctx)
	if err != nil {
		utils.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, "couldn't unsubscribe email")
		return
	}
	utils.Logger.Info("email unsubsribed")
	c.JSON(http.StatusOK, "sucess")
}
