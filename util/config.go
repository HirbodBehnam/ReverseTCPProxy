package util

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
)

// ParseConfig will parse the config file into given structure pointer
func ParseConfig(config any) {
	// Set the name
	configName := "config.json"
	if len(os.Args) == 1 {
		configName = os.Args[0]
	}
	log.WithField("name", configName).Trace("loading config")
	// Open the file
	configFile, err := os.Open(configName)
	if err != nil {
		log.WithError(err).Fatal("cannot open config file")
	}
	// Parse
	err = json.NewDecoder(configFile).Decode(config)
	if err != nil {
		log.WithError(err).Fatal("cannot parse config file")
	}
	_ = configFile.Close()
	log.WithField("config", config).Trace("loaded config")
}
