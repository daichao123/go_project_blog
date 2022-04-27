package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_project_blog/controller"
)

func TestRouter() {
	fmt.Printf("test router")
}

func Start() {
	engine := gin.Default()
	engine.POST("/register", controller.Register)
	engine.Run()
}
