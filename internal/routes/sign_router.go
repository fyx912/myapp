package routes

import (
	handler "myapp/internal/handler"
	"myapp/internal/repositories"
	"myapp/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignRrouter(router *gin.Engine, db *gorm.DB) {
	signRepo := repositories.NewSignRepository(db)
	signSerive := service.NewSignSerive(signRepo)
	signHandler := handler.NewSignHandler(signSerive)

	router.POST("/signIn", signHandler.SignIn)
	router.POST("/signOut", signHandler.SignOut)
}
