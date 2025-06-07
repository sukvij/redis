package controller

import (
	"fmt"
	"net/http"
	"strconv"
	userModel "vijju/user/model"
	"vijju/user/service"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Controller struct {
	Db          *gorm.DB
	RedisClient *redis.Client
}

func NewController(db *gorm.DB, redisClient *redis.Client) *Controller {
	return &Controller{Db: db, RedisClient: redisClient}
}

func (controller *Controller) UserAPI(app *gin.Engine) {
	app.GET("/users", controller.getAllUsers)
	app.GET("/users/:id", controller.getUserById)
	app.POST("/users", controller.createUser)
}

func (controller *Controller) getAllUsers(ctx *gin.Context) {
	service := service.NewService(nil, controller.Db, controller.RedisClient)
	result, err := service.GetAllUsers()
	if err != nil {
		ctx.JSON(-1, err)
		return
	} else {
		ctx.JSON(200, result)
	}
}

func (controller *Controller) getUserById(ctx *gin.Context) {
	id, err1 := strconv.Atoi(ctx.Param("id"))
	if err1 != nil {
		fmt.Println("string to int convert error err ", err1)
		return
	}
	service := service.NewService(&userModel.User{ID: uint(id)}, controller.Db, controller.RedisClient)
	result, err := service.GetUserById()
	if err != nil {
		ctx.JSON(-1, err)
		return
	} else {
		ctx.JSON(200, result)
	}
}

func (controller *Controller) createUser(ctx *gin.Context) {
	user := &userModel.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// controller.User = user
	service := service.NewService(user, controller.Db, controller.RedisClient)
	result, err := service.CreateUser()
	if err != nil {
		ctx.JSON(-1, err)
		return
	} else {
		ctx.JSON(200, result)
	}
}
