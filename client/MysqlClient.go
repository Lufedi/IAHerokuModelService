package client

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ModelsService/model"
)

const (
	username = "root"
	password = "secret"
	hostname = "127.0.0.1:3306"
	dbname   = "iaheroku"
)

func database() *gorm.DB {
	dsn := "root:secret@tcp(127.0.0.1:3306)/iaheroku?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Errorf("Error opening connecction", err)
	}else {
		fmt.Println("database created")
	}
	return db
}

func migrate() *gorm.DB {
	db := database()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Model{})

	/*db.Create(&model.User{
		Rol: "user",
		Email: "myemail@gmail.com",
	})*/
	return db
}

func dns() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

var Db = migrate()