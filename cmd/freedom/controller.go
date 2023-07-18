package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func controllerEndpoint(w http.ResponseWriter, r *http.Request) {
	logger := log.WithField("remote", r.RemoteAddr)
	logger.Debug("controller called")
	// Parse the form and get the id
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.WithError(err).Warn("cannot parse form")
		return
	}
	id := r.Form.Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		logger.Warn("empty body")
		return
	}
	// Connect to remote host
	w.WriteHeader(http.StatusNoContent)
	go initiateReverseProxy(id)
}
