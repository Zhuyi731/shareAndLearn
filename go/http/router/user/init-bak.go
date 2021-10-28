package user

import (
	"github.com/gin-gonic/gin"
	userController "http/controller/user"
)

func InitRouter(v *gin.RouterGroup) {
	userRouter := v.Group("/user")
	{
		userRouter.POST("/", userController.AddUser)
		userRouter.GET("/", userController.GetUser)
	}
}
