package repositories

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetUsers() ([]models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) error
}
type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (u *userRepo) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) CreateUser(user *models.User) (*models.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepo) UpdateUser(user *models.User) (*models.User, error) {
	if err := u.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepo) DeleteUser(user *models.User) error {
	if err := u.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
