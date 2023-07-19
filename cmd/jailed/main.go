package main

import "ReverseTCPProxy/util"

func main() {
	util.ParseConfig(&config)
	pendingConnections = util.NewPendingConnections()
	go listenForLocal()
	go listenForRemote()
	select {} // wait forever...
}
