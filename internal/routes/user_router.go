package routes

import (
	handler "myapp/internal/handler"
	"myapp/internal/repositories"
	"myapp/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUserRouter(router *gin.Engine, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	user := router.Group("/user")
	{
		user.GET("", userHandler.GetUser)
		user.GET("list", userHandler.GetUserArray)
		user.GET(":uid", userHandler.GetUserByID)
		user.PUT(":uid", userHandler.UpdateByID)
		user.DELETE(":uid", userHandler.DeleteUserByID)
		user.POST("", userHandler.SaveUser)
	}
}
