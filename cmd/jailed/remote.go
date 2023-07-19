package main

import (
	log "github.com/sirupsen/logrus"
	"net"
	"regexp"
	"strconv"
	"time"
)

// regex pattern of the first remote packet
var firstPacketRegex = regexp.MustCompile("GET /(\\d+) HTTP/1.1\r\n\r\n")

// listenForRemote will listen for remote connections
func listenForRemote() {
	// Create a listener
	listener, err := net.Listen("tcp", config.FreedomListenAddress)
	if err != nil {
		log.WithError(err).Fatal("cannot listen for freedom connections")
	}
	// Serve each TCP connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.WithError(err).Error("cannot accept freedom connection")
		}
		// Maybe on same thread is a good idea because we are only going to read one message from it
		handleRemoteConnection(conn)
	}
}

func handleRemoteConnection(remoteConn net.Conn) {
	logger := log.WithField("remote", remoteConn.RemoteAddr())
	logger.Debug("new remote connection!")
	// Read first packet
	buffer := make([]byte, 64)
	n, err := remoteConn.Read(buffer)
	if err != nil {
		logger.WithError(err).Warning("cannot read first packet of remote connection")
		_ = remoteConn.Close()
		return
	}
	// Otherwise parse it
	id, firstPacketOk := parseRemoteFirstPacket(buffer[:n])
	if !firstPacketOk {
		logger.WithField("first_packet", string(buffer[:n])).Warning("invalid first packet")
		_ = remoteConn.Close()
		return
	}
	logger = logger.WithField("id", id)
	logger.Debug("got first packet")
	// Send it to local connection
	timer := time.NewTimer(config.RemoteWaitTimeout)
	select {
	case remoteConnectionPool <- remoteConn:
		logger.Debug("matched remote with local")
		timer.Stop()
	case <-timer.C:
		_ = remoteConn.Close()
		logger.Error("cannot match remote with local")
	}
}

// parseRemoteFirstPacket parses the first packet of a remote connection.
// First returned number is the ID of connection and the second one is a bool which indicates that if
// the parse was successful or not.
func parseRemoteFirstPacket(firstPacket []byte) (uint32, bool) {
	matches := firstPacketRegex.FindSubmatch(firstPacket)
	if matches == nil { // no match
		return 0, false
	}
	// Parse the id
	number, err := strconv.ParseUint(string(matches[1]), 10, 32)
	return uint32(number), err == nil
}
