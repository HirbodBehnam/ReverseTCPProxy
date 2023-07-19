package util

import (
	"net"
	"sync"
)

// PendingConnections is a map for pending connections which are waiting for their remote connection to be
// established. The point is, the map goes from an ID to a channel which accepts a net connection.
type PendingConnections struct {
	connections map[uint32]chan net.Conn
	mu          sync.Mutex
}

func NewPendingConnections() *PendingConnections {
	return &PendingConnections{connections: make(map[uint32]chan net.Conn)}
}

// MatchConnection will send a remote connection to a local one to be matched and then proxied.
// Returns false if there is no corresponding local connection, otherwise true.
func (c *PendingConnections) MatchConnection(id uint32, conn net.Conn) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	channel := c.connections[id]
	if channel == nil { // not found
		return false
	} else {
		channel <- conn
		return true
	}
}

// RegisterLocalConnection will create a channel pending for a new connection + an id for the connection.
func (c *PendingConnections) RegisterLocalConnection() (uint32, <-chan net.Conn) {
	// Create a channel and ID
	id := RandomID()
	channel := make(chan net.Conn, 1)
	// Add them to map
	c.mu.Lock()
	c.connections[id] = channel
	c.mu.Unlock()
	// Return
	return id, channel
}

// DitchConnection will ditch a pending local connection
func (c *PendingConnections) DitchConnection(id uint32) {
	// At first get the channel and delete the id from map
	c.mu.Lock()
	channel := c.connections[id]
	delete(c.connections, id)
	c.mu.Unlock()
	// Now check if the channel is empty or not. If it's not empty, it means that the server is connection
	// in the instance which client is timed out. In this case, we shall close the remote connection.
	if channel != nil {
		select {
		case conn, _ := <-channel:
			_ = conn.Close()
		default:
			// nothing! Channel is empty
		}
	}
}
