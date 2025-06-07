package service

import (
	"encoding/json"
	"fmt"
	userModel "vijju/user/model"
	"vijju/user/repository"

	redisService "vijju/redis"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Service struct {
	User        *userModel.User
	Db          *gorm.DB
	RedisClient *redis.Client
}

func NewService(user *userModel.User, db *gorm.DB, redisClient *redis.Client) *Service {
	return &Service{User: user, Db: db, RedisClient: redisClient}
}

func (service *Service) GetAllUsers() (*[]userModel.User, error) {
	// find result from redis and return
	repository := repository.NewRepository(service.User, service.Db, service.RedisClient)
	result, err := repository.GetAllUsers()
	return result, err
}

func (service *Service) CreateUser() (*userModel.User, error) {
	repository := repository.NewRepository(service.User, service.Db, service.RedisClient)
	result, err := repository.CreateUser()
	if err == nil {
		key := fmt.Sprintf("user:%d", service.User.ID)
		byteData, _ := json.Marshal(service.User)
		value := string(byteData)
		fmt.Println("key and val is ", key, value)
		if errs := redisService.SetValue(service.RedisClient, key, value); errs != nil {
			fmt.Println("create user --> cannot store user creation - errs is ", errs)
		}
	}
	return result, err
}
func (service *Service) GetUserById() (*userModel.User, error) {
	// find result from redis if exist other wise get from database and save in the redis
	key := fmt.Sprintf("user:%d", service.User.ID)
	redisRes, redErr := redisService.GetValue(service.RedisClient, key)
	if redErr != nil {
		fmt.Println("redis get user error - ", redErr)
	} else {
		if redisRes != "" {
			json.Unmarshal([]byte(redisRes), service.User)
			return service.User, nil
		}
	}

	// mean result not found in redis so go to database and fetch result from database
	fmt.Println("we enterned into database and fetch result from database..")
	repository := repository.NewRepository(service.User, service.Db, service.RedisClient)
	result, err := repository.GetUserById()
	if err == nil {
		key := fmt.Sprintf("user:%d", service.User.ID)
		byteData, _ := json.Marshal(service.User)
		value := string(byteData)
		fmt.Println("key and val is ", key, value)
		if errs := redisService.SetValue(service.RedisClient, key, value); errs != nil {
			fmt.Println("getuser by id --> cannot store user creation - errs is ", errs)
		}
	}
	return result, err
}
