package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/himanshu07070/newsletter/Handlers"
	input "github.com/himanshu07070/newsletter/Input"
	utils "github.com/himanshu07070/newsletter/Utils"
)

func Middleware(dataChannel chan input.MetaData) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("channel", dataChannel)
		c.Next()
	}
}
func Router(dataChannel chan input.MetaData) *gin.Engine {
	utils.Logger.Info("Inside router...")
	router := gin.Default()
	router.Use(Middleware(dataChannel))
	router.POST("/unsubscribe", handlers.UnsubscribeEmails)
	router.POST("/subscribe", handlers.SaveEmails)
	router.POST("/newsletter", handlers.Newsletter)
	return router
}
