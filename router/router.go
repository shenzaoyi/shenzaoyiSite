package router

import (
	"be/controller"
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	//	版本
	v1 := e.Group("/v1")

	userGroup := v1.Group("/user")                 //	用户组
	userGroup.POST("/login", controller.LoginByUP) //	用户登录
	userGroup.POST("/register", controller.Register)
}
