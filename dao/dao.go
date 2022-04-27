package dao

import (
	"go_project_blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Manager interface {
	Register(user *model.User)
	Update(id int, user *model.User)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

// Register 注册用户
func (manager manager) Register(user *model.User) {
	manager.db.Create(user)
}

// Update 更新指定字段
func (manager manager) Update(id int, user *model.User) {
	//manager.db.Debug().Where("id = ? ", id).Update("username", user.Username)
	update := manager.db.Model(user).Where("id = ?", id).Update("username", user.Username)
	log.Fatal(update.RowsAffected)
}

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("链接数据库错误")
	}
	Mgr = &manager{db: db}
	err1 := db.AutoMigrate(&model.User{})
	if err1 != nil {
		log.Fatal(err1)
	}
}
