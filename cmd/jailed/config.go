package main

import "time"

var config struct {
	// On what interface/port should we listen for local incoming connection?
	LocalListenAddress string `json:"local_listen_address"`
	// On what interface/port should we expect our freedom server to connect to us?
	FreedomListenAddress string `json:"freedom_listen_address"`
	// What is the control endpoint address?
	ControlEndpoint string `json:"control_endpoint"`
	// Timeout which we should wait before killing a local connection.
	// Unit is milliseconds.
	RemoteWaitTimeout time.Duration `json:"remote_wait_timeout"`
}
