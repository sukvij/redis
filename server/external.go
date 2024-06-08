package server

import (
	"fmt"
	"vijju/router"
)

var (
	externalShutdown       = make(chan struct{})
	externalShutdownFinish = make(chan struct{})
)

func buildExternalRouter() *router.Router {

	router := router.CreateRouter()

	addUserRouter(router)

	return router
}

func startExternalServer() {
	fmt.Println("starting external server")
	defer doApiRecovery()

	// logging info()

	externalRouter := buildExternalRouter()
	externalRouter.Run(":8080")
}

func stopExternalServer() <-chan struct{} {
	return externalShutdown
}
