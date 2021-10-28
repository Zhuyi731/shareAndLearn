package userController

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("xxx2")
}

type getUserParam struct {
	Id int `form:"id"`
}
//
func GetUser(ctx *gin.Context) {
	param := getUserParam{}
	ctx.ShouldBindQuery(&param)

	result := users[param.Id]
	ctx.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}
