package router

import (
	"github.com/gin-gonic/gin"
	"project/router/Mongo"
	"project/router/line"
)

func NewLineRoute(group *gin.RouterGroup) *gin.RouterGroup {
	group.POST("", line.StoreMessage)
	group.POST("/push", line.PushMessage)
	return group
}
func NewUserRoute(group *gin.RouterGroup) *gin.RouterGroup {
	group.GET("/:userid/messages", Mongo.GetUserMessage)
	return group
}
