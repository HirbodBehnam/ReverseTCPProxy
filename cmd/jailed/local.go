package main

import (
	"net"
	"sync"
)

// pendingConnections is a map for pending connections which are waiting for their remote connection to be
// established. The point is, the map goes from an ID to a channel which accepts a net connection.
// TODO: move these to a struct
var pendingConnections = make(map[uint32]chan<- *net.Conn)
var pendingConnectionsMutex sync.Mutex

func listenForLocal() {

}
