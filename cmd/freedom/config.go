package main

var config struct {
	// On what interface/port should we listen for control commands.
	// This should probably be reversed proxied through nginx or caddy.
	ControlListen string
	// Address of the computer which has jailed internet connection
	JailedComputerAddress string
	// After we have established a reverse proxy, where should we proxy the traffic to?
	DestinationAddress string
}
