package router

import (
	handler "app1/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	router := gin.Default()

	v1 := router.Group("/user")
	{
		v1.GET("/login", handler.Login)
		v1.POST("/register", handler.Register)
	}

	v2 := router.Group("/msg")
	{
		v2.GET("/getMsgs", handler.GetMsgs)
		v2.POST("/add", handler.AddMsg)
	}

	return router
}
