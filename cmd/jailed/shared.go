package main

import (
	"net"
)

// Remote connection pool
var remoteConnectionPool = make(chan net.Conn)
