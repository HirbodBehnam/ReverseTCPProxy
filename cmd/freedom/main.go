package main

import (
	"ReverseTCPProxy/util"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetLevel(log.TraceLevel)
	util.ParseConfig(&config)
	http.HandleFunc("/", controllerEndpoint)
	err := http.ListenAndServe(config.ControlListen, nil)
	if err != nil {
		log.WithError(err).Fatal("cannot listen for controller")
	}
}
