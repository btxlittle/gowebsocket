// Package routers 路由
package routers

import (
	"github.com/gin-gonic/gin"

	"g2git.hdyouxi.com/xiandi/game-socket-server/controllers/home"
	"g2git.hdyouxi.com/xiandi/game-socket-server/controllers/systems"
	"g2git.hdyouxi.com/xiandi/game-socket-server/controllers/user"
)

// Init http 接口路由
func Init(router *gin.Engine) {
	router.LoadHTMLGlob("views/**/*")

	// 用户组
	userRouter := router.Group("/user")
	{
		userRouter.GET("/list", user.List)
		userRouter.GET("/online", user.Online)
		userRouter.POST("/sendMessage", user.SendMessage)
		userRouter.POST("/sendMessageAll", user.SendMessageAll)
	}

	// 系统
	systemRouter := router.Group("/system")
	{
		systemRouter.GET("/state", systems.Status)
	}

	// home
	homeRouter := router.Group("/home")
	{
		homeRouter.GET("/index", home.Index)
	}
}
