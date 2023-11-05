package handler

import (
	"encoding/json"
	"log"
	"myapp/internal/models"
	"myapp/internal/service"
	result "myapp/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	uidStr := c.Param("uid")

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Err.ToString())
		return
	}
	log.Println("uid:", uid)

	// 获取用户
	existingUser, _ := uh.userService.GetUserByID(uid)

	// c.JSON(http.StatusOK, existingUser)
	c.JSON(http.StatusOK, result.OK.WithData(existingUser))
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	account := c.Param("account")
	name := c.Param("name")
	user := &models.User{}
	user.Account = account
	user.Name = name

	// 获取用户
	existingUser, _ := uh.userService.GetUser(user)

	// c.JSON(http.StatusOK, existingUser)
	c.JSON(http.StatusOK, result.OK.WithData(existingUser))
}

func (uh *UserHandler) GetUserArray(c *gin.Context) {
	account := c.Param("account")
	name := c.Param("name")
	user := &models.User{}
	user.Account = account
	user.Name = name

	userJson, _ := json.Marshal(user)
	log.Println("requset data: v%", string(userJson))

	// 获取用户
	users, _ := uh.userService.GetUserArray(user)

	// c.JSON(http.StatusOK, existingUser)
	c.JSON(http.StatusOK, result.OK.WithData(users))
}

func (uh *UserHandler) UpdateByID(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, err := strconv.Atoi(uidStr)
	log.Println("uid:", uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Err.ToString())
		return
	}
	// 获取用户
	user, _ := uh.userService.GetUserByID(uid)
	// user := &models.User{}
	user.Uid = uid

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, result.Err.ToString())
		return
	}
	// 获取用户
	if err := uh.userService.UpdateUserByID(user); err != nil {
		c.JSON(http.StatusBadRequest, result.Err.ToString())
		return
	}

	// c.JSON(http.StatusOK, existingUser)
	c.JSON(http.StatusOK, result.OK)
}

func (uh *UserHandler) SaveUser(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, result.Err.ToString())
		return
	}
	json, _ := json.Marshal(user)
	log.Printf("request data:{} %v,", string(json))

	// 获取用户
	// user, _ := uh.userService.GetUserByID(user.Uid)

	// 获取用户
	if err := uh.userService.SaveUser(user); err != nil {
		c.JSON(http.StatusBadRequest, result.Err.ToString())
		return
	}

	// c.JSON(http.StatusOK, existingUser)
	c.JSON(http.StatusOK, result.OK)
}

func (uh *UserHandler) DeleteUserByID(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, err := strconv.Atoi(uidStr)
	log.Println("uid:", uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Err.ToString())
		return
	}
	if err := uh.userService.DeleteUserByID(uid); err != nil {
		c.JSON(http.StatusBadRequest, result.Err.ToString())
		return
	}
	c.JSON(http.StatusOK, result.OK)
}
