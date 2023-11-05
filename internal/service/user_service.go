package service

import (
	"encoding/json"
	"myapp/internal/models"
	"myapp/internal/repositories"
	"myapp/internal/utils"
	stringUtils "myapp/internal/utils/stringUtils"
	"time"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	// 创建用户服务实例，并将userRepo赋值给userRepo字段
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) GetUserByID(uid int) (*models.User, error) {
	return us.userRepo.GetByID(uid)
}

func (us *UserService) GetUser(user *models.User) (*models.User, error) {
	return us.userRepo.GetUser(user)
}

func (us *UserService) GetUserArray(user *models.User) ([]*models.User, error) {
	return us.userRepo.GetUserArray(user)
}

func (us *UserService) UpdateUserByID(user *models.User) error {
	user.UpdateTime = utils.ToLocalDateTime(time.Now())
	return us.userRepo.Update(user)
}

func (us *UserService) SaveUser(user *models.User) error {
	if stringUtils.IsBlank(user.Account) {
		panic("Account 不能为空") // 抛出错误
	}

	if user.Status == 0 {
		user.Status = 1
	}
	jsonData, _ := json.Marshal(user)
	user.Salt = utils.MD5(string(jsonData))

	user.CreateTime = utils.ToLocalDateTime(time.Now())
	user.UpdateTime = user.CreateTime
	return us.userRepo.Create(user)
}

func (us *UserService) DeleteUserByID(uid int) error {
	user := &models.User{}
	user.Uid = uid
	return us.userRepo.Delete(user)
}
