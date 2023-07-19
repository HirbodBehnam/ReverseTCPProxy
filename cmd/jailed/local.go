package main

import (
	"ReverseTCPProxy/util"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// listenForLocal will listen for local connections and signals the remote server to establish a connection from its side
// to us.
func listenForLocal() {
	// Create a TCP listener
	listener, err := net.Listen("tcp", config.LocalListenAddress)
	if err != nil {
		log.WithError(err).Fatal("cannot listen for local connections")
	}
	// Serve each TCP connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.WithError(err).Error("cannot accept local connection")
		}
		go handleLocalConnection(conn)
	}
}

// handleLocalConnection will handle a newly created local connection
func handleLocalConnection(localConn net.Conn) {
	defer localConn.Close()
	logger := log.WithField("local", localConn.LocalAddr())
	// At first, we should notify the remote host that we need a new connection
	id := util.RandomID()
	logger = logger.WithField("id", id)
	logger.Debug("sending message to control")
	resp, err := http.PostForm(config.ControlEndpoint, url.Values{"id": []string{strconv.FormatUint(uint64(id), 10)}})
	if err != nil {
		logger.WithError(err).Error("cannot send request to control")
		return
	}
	_ = resp.Body.Close()
	// Check status code
	if resp.StatusCode/100 != 2 {
		logger.WithField("status_code", resp.StatusCode).WithField("status", resp.Status).Error("not 2xx status code from control")
		return
	}
	// Now wait for connection from freedom
	var freedomConn net.Conn
	timer := time.NewTimer(config.RemoteWaitTimeout)
	select { // either ...
	case freedomConn = <-remoteConnectionPool: // we get the connection from remote host
		// avoid leaks
		timer.Stop()
		break
	case <-timer.C: // or we have a timeout :(
		logger.Warning("timeout on waiting for remote host")
		return // return from method. This closes the local connection
	}
	// From now on, we have the remote connection!
	defer freedomConn.Close()
	logger = logger.WithField("remote", freedomConn.RemoteAddr())
	logger.Debug("match!")
	// Now simply proxy everything!
	err = util.ProxyConnection(freedomConn, localConn)
	if err != nil {
		logger.WithError(err).Error("cannot proxy")
	}
	logger.Debug("proxy done")
}
