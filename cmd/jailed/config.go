package main

var config struct {
	// On what interface/port should we listen for local incoming connection?
	LocalListenAddress string
	// On what interface/port should we expect our freedom server to connect to us?
	FreedomListenAddress string
	// What is the control endpoint address?
	ControlEndpoint string
}
