package main

import (
	"ReverseTCPProxy/util"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	log.SetLevel(log.TraceLevel)
	util.ParseConfig(&config)
	// Set timeout
	if config.RemoteWaitTimeout == 0 {
		config.RemoteWaitTimeout = 10_000 // 10 sec default wait time
	}
	config.RemoteWaitTimeout *= time.Millisecond
	log.WithField("remote_timeout", config.RemoteWaitTimeout).Trace("timeout set")
	// Start listeners
	pendingConnections = util.NewPendingConnections()
	go listenForLocal()
	go listenForRemote()
	select {} // wait forever...
}
