package main

import (
	"ReverseTCPProxy/util"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.TraceLevel)
	util.ParseConfig(&config)
	if config.RemoteWaitTimeout == 0 {
		config.RemoteWaitTimeout = 10_000 // 10 sec default wait time
	}
	pendingConnections = util.NewPendingConnections()
	go listenForLocal()
	go listenForRemote()
	select {} // wait forever...
}
