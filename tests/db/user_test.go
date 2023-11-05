package db

import (
	"log"
	"myapp/internal/models"
	"myapp/internal/utils"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SelectUser() {
	// 加载配置文件
	config, err := utils.LoadConfig("configs/db.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 连接数据库
	db, err := gorm.Open("mysql", config.Database.UserName+":"+config.Database.Password+"@tcp("+config.Database.Host+":"+strconv.Itoa(config.Database.Port)+")/"+config.Database.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	defer db.Close()

	user := models.User
	db.First(user)

}
