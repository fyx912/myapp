package routes

import (
	"myapp/internal/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	router.Use(middlewares.PanicHandler)

	SignRrouter(router, db)

	SetUserRouter(router, db)

	v1 := router.Group("/v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "你好，世界！"})
		})
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "hello world"})
		})
	}

}
