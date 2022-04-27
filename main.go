package main

import (
	"go_project_blog/dao"
	"go_project_blog/model"
)

func main() {
	//controller.TestController()
	user := model.User{
		Username: "update username",
		Password: "12345678",
	}
	dao.Mgr.Update(1, &user)
	//dao.Mgr.Register(&user)
}
