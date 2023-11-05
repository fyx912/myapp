package repositories

import (
	"myapp/internal/models"
	dto "myapp/internal/models/dto"

	"gorm.io/gorm"
)

type SignRepository struct {
	db *gorm.DB
}

func NewSignRepository(db *gorm.DB) *SignRepository {
	return &SignRepository{
		db: db,
	}
}

func (sr *SignRepository) SignIn(signDTO *dto.SignInDTO) (*models.User, error) {
	user := &models.User{}
	err := sr.db.Where("account = ? or phone_number=?  or email=?", signDTO.Account, signDTO.Account, signDTO.Account).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
