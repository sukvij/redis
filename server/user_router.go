package server

import (
	"vijju/api"
	"vijju/router"
)

func expostUserEndpoints(router *router.Router) {
	router.RouterGroup.GET("/getAllUsers", api.GetAllUsers)
}

func addUserRouter(router *router.Router) {
	router.RouterGroup = router.App.Group("/user/v1")
	{
		// we can make changes in the router
		router.BasePath += "/user/v1"
		router.Title += "user apis"
		expostUserEndpoints(router)
	}
}
