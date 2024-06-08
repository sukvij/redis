package server

const (
	logTag = "server"
)

// start the external server
func Start() {
	go startExternalServer()
}

// stop all the servers

func Stop() <-chan struct{} {
	return WaitChannels(stopExternalServer())
}

func doApiRecovery() {}

func WaitChannels(chs ...<-chan struct{}) <-chan struct{} {
	ret := make(chan struct{})
	go func() {
		for _, ch := range chs {
			<-ch
		}
		close(ret)
	}()
	return ret
}
