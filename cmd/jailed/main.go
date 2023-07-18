package main

import "ReverseTCPProxy/util"

func main() {
	util.ParseConfig(&config)
	go listenForLocal()
	go listenForRemote()
	select {} // wait forever...
}
