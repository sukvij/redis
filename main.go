package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"vijju/server"
)

// const shutDownTimeout = 10 * time.Second

func main() {
	fmt.Println("running main.go")
	server.Start()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
	// select {
	// // case <-concurrenctutils.WaitChannels(server.Stop()):
	// case <-time.After(shutDownTimeout):
	// case
	// }
}

func Shutdown() <-chan struct{} {
	// shutdown closes the connection to the sattusD server
	return nil
}
