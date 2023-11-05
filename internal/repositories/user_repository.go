package repositories

import (
	"myapp/internal/models"
	stringUtils "myapp/internal/utils/stringUtils"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) Update(user *models.User) error {
	return ur.db.Save(user).Error
}

func (ur *UserRepository) Delete(user *models.User) error {
	return ur.db.Delete(user).Error
}

func (ur *UserRepository) GetByID(uid int) (*models.User, error) {
	var user models.User
	err := ur.db.First(&user, uid).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetUser(user *models.User) (*models.User, error) {
	var userData models.User
	query := ur.db.Where("status=1")
	if stringUtils.IsNotBlank(user.Account) {
		query = query.Where("account=?", user.Account)
	}
	if stringUtils.IsNotBlank(user.Name) {
		query = query.Where("name = ?", user.Name)
	}
	err := query.Find(&userData).Error
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

func (ur *UserRepository) GetUserArray(user *models.User) ([]*models.User, error) {
	var users []*models.User

	query := ur.db.Where("1=1")
	if stringUtils.IsNotBlank(user.Account) {
		query = query.Where("account=?", user.Account)
	}
	if stringUtils.IsNotBlank(user.Name) {
		query = query.Where("name = ?", user.Name)
	}
	result := query.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
