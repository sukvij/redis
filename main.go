package main

import (
	"vijju/database"
	"vijju/redis"
	userController "vijju/user/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()
	redisClient := redis.NewRedisClient()
	app := gin.Default()
	(userController.NewController(db, redisClient)).UserAPI(app)
	if err := app.Run(":8080"); err != nil {
		panic("server is not runnig...")
	}
}
