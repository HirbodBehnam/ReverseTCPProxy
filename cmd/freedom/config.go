package main

var config struct {
	// On what interface/port should we listen for control commands.
	// This should probably be reversed proxied through nginx or caddy.
	ControlListen string `json:"control_listen"`
	// Address of the computer which has jailed internet connection
	JailedComputerAddress string `json:"jailed_computer_address"`
	// After we have established a reverse proxy, where should we proxy the traffic to?
	DestinationAddress string `json:"destination_address"`
}
