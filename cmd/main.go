package main

import (
	"log"
	"myapp/internal/routes"
	"myapp/internal/utils"
	"net/http"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	// 加载配置文件
	config, err := utils.LoadConfig("configs/db.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db := connectDatabase(config)
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("database is closed, case: %v", err)
	}
	defer sqlDB.Close()

	appRouter(db)
}

func connectDatabase(config *utils.Config) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	dsn := config.Database.UserName + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + strconv.Itoa(config.Database.Port) + ")/" + config.Database.DBName + "?charset=utf8&parseTime=True&loc=Local"
	log.Println(dsn)
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func appRouter(db *gorm.DB) {
	//1.创建路由
	router := gin.Default()
	routes.SetupRoutes(router, db)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
