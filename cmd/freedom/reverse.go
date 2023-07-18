package main

import (
	"ReverseTCPProxy/util"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

// initiateReverseProxy will initiate a reverse proxy with specified ID
func initiateReverseProxy(id string) {
	logger := log.WithField("id", id)
	// Dial jailed endpoint
	jailedComputer, err := net.Dial("tcp", config.JailedComputerAddress)
	if err != nil {
		logger.WithError(err).Error("cannot dial jailed computer")
		return
	}
	defer jailedComputer.Close()
	// Send the initial message
	_, err = fmt.Fprintf(jailedComputer, "GET /%s HTTP/1.1\n\r\n\r", id)
	if err != nil {
		logger.WithError(err).Error("cannot send hello message to jailed computer")
		return
	}
	// Dial our endpoint
	destination, err := net.Dial("tcp", config.DestinationAddress)
	if err != nil {
		logger.WithError(err).Error("cannot dial destination")
		return
	}
	defer destination.Close()
	// Proxy everything
	err = util.ProxyConnection(jailedComputer, destination)
	if err != nil {
		log.WithError(err).Warn("cannot proxy connection")
	}
}
