package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	App             *gin.Engine
	RouterGroup     *gin.RouterGroup
	Title           string
	Description     string
	Version         string
	BasePath        string
	ReadTimeOut     time.Duration
	WriteTimeOut    time.Duration
	ShutDownTimeOut time.Duration
}

func CreateRouter() *Router {
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://127.0.0.1:8080", "http://localhost:8081", "http://0.0.0.0:8080"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router := &Router{
		App:             app,
		Title:           "App ",
		Description:     "app_rest_apis_",
		Version:         "1.1",
		BasePath:        "",
		ReadTimeOut:     10 * time.Second,
		WriteTimeOut:    10 * time.Second,
		ShutDownTimeOut: 3 * time.Second,
	}
	return router
}

func (router *Router) Run(port string) {
	router.App.Run(port)
}
