package repository

import (
	userModel "vijju/user/model"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Repository struct {
	User        *userModel.User
	Db          *gorm.DB
	RedisClient *redis.Client
}

func NewRepository(user *userModel.User, db *gorm.DB, redisClient *redis.Client) *Repository {
	return &Repository{User: user, Db: db, RedisClient: redisClient}
}

func (repository *Repository) GetAllUsers() (*[]userModel.User, error) {
	var users *[]userModel.User
	if err := repository.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *Repository) CreateUser() (*userModel.User, error) {
	// var users *[]userModel.User
	if err := repository.Db.Create(repository.User).Error; err != nil {
		return nil, err
	}
	return repository.User, nil
}

func (repository *Repository) GetUserById() (*userModel.User, error) {
	// var users *[]userModel.User
	if err := repository.Db.First(repository.User.ID).Error; err != nil {
		return nil, err
	}
	return repository.User, nil
}
