package service

import (
	"log"
	"myapp/internal/models/dto"
	"myapp/internal/repositories"
	result "myapp/internal/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type SignService struct {
	signRepo *repositories.SignRepository
}

func NewSignSerive(signRepo *repositories.SignRepository) *SignService {
	return &SignService{
		signRepo: signRepo,
	}
}

func (ss *SignService) SignIn(signDTO *dto.SignInDTO) (result.Response, int) {
	httpStatus := http.StatusBadRequest
	user, err := ss.signRepo.SignIn(signDTO)
	if err != nil {
		log.Printf("SignIn failed : {} %v", err)
		return result.ErrUserAccountOrPasswd.Failed(), httpStatus
	}
	// 判断用户名密码是否正确
	if signDTO.Account != user.Account || signDTO.Passwd != user.Passwd {
		log.Println(result.ErrUserAccountOrPasswd.ToString())
		return result.ErrUserAccountOrPasswd.Failed(), httpStatus
	}
	//redis
	uuid := uuid.New()
	uuidString := strings.ToUpper(strings.ReplaceAll(uuid.String(), "-", ""))
	return result.OK.WithData(uuidString), http.StatusOK
}
