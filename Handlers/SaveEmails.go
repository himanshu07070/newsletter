package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	database "github.com/himanshu07070/newsletter/Database"
	input "github.com/himanshu07070/newsletter/Input"
	utils "github.com/himanshu07070/newsletter/Utils"
)

func SaveEmails(c *gin.Context) {
	utils.Logger.Info("Inside SaveEmails...")
	ctx := c.Request.Context()
	var userInput input.User
	err := c.BindJSON(&userInput)
	if err != nil {
		utils.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, "invalid input")
		return
	}
	if userInput.Email == "" || !strings.Contains(userInput.Email, "@") {
		utils.Logger.Info("Invalid email Id")
		c.JSON(http.StatusBadRequest, "invalid email")
		return
	}
	if !utils.ValidateEmailAddress(userInput.Email) || !utils.IsEmailValid(userInput.Email) {
		utils.Logger.Info("Invalid email Id")
		c.JSON(http.StatusBadRequest, "invalid email")
		return
	}

	_, exist := database.CheckUserEmail(userInput.Email, ctx)
	if exist {
		utils.Logger.Info("user already exist")
		c.JSON(http.StatusBadRequest, "user already exist")
		return
	}
	var user database.User
	user.Email = userInput.Email
	user.Tag = userInput.Tag
	user.Subscribe = true
	if err := database.InsertUserDetail(&user, ctx); err != nil {
		utils.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, "failure")
		return
	}
	utils.Logger.Info("record inserted")
	c.JSON(http.StatusOK, "sucess")
}
