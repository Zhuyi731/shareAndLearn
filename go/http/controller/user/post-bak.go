package userController

import (
	"fmt"
)

type User struct {
	Name     string
	Account  string
	Password string
}

var users map[int]User

var id int

func init() {
	id = 0
	users = make(map[int]User)
	fmt.Println("xxx1")
}

type userParam struct {
	Name     string `form:"name"`
	Account  string `form:"account"`
	Password string `form:"password"`
}
//
//func AddUser(ctx *gin.Context) {
//	param := userParam{}
//	ctx.ShouldBind(&param)
//
//	 users[id] = User{
//		Name:     param.Name,
//		Account:  param.Account,
//		Password: param.Password,
//	}
//	id++
//
//	ctx.JSON(200, gin.H{
//		"code": 0,
//		"data": users[id-1],
//	})
//}
