package handler

import (
	"myapp/internal/models/dto"
	"myapp/internal/service"
	result "myapp/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignHandler struct {
	signSerive *service.SignService
}

func NewSignHandler(signSerive *service.SignService) *SignHandler {
	return &SignHandler{
		signSerive: signSerive,
	}
}

func (handler *SignHandler) SignIn(c *gin.Context) {
	// 声明接收的变量
	var signInDTO dto.SignInDTO
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&signInDTO); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, result.ErrParam.ToString())
		return
	}

	data, httpStatus := handler.signSerive.SignIn(&signInDTO)
	if httpStatus != http.StatusOK {
		c.JSON(httpStatus, data)
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "status": "200", "message": "success", "authorization": data.Data})
	}
}

func (handler *SignHandler) SignOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "status": "200", "message": "success"})
}
