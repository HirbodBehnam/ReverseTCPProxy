package main

import "ReverseTCPProxy/util"

// List of pending local connections, waiting for a remote connection to be matched with
var pendingConnections *util.PendingConnections
