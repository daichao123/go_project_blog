package model

import (
	"fmt"
	"gorm.io/gorm"
)

func TestModel() {
	fmt.Printf("test")
}

// User 用户模型
//模拟mysql User表
//蛇型复数格式
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
