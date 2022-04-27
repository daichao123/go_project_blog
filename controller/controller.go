package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_project_blog/dao"
	"go_project_blog/model"
)

func TestController() {
	fmt.Printf("test controller")
}

// Register 模拟用户注册
func Register(ctx *gin.Context) {
	username := ctx.DefaultPostForm("username", "")
	password := ctx.DefaultPostForm("password", "")
	userFomat := model.User{
		Username: username,
		Password: password,
	}
	dao.Mgr.Register(&userFomat)
}
