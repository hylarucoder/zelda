package main

import (
	"github.com/sirupsen/logrus"
	"github.com/twocucao/zelda/app/www"
	"github.com/twocucao/zelda/library/environments"
	"os"
)

type contextKey int

const (
	// ServiceName is how this app is identified in logs and error handlers
	ServiceName          string     = "faraday"
	userID               contextKey = iota // Used for gorilla/context
	userSudo             contextKey = iota // Used for gorilla/context
	requestAuthenticated contextKey = iota // Used for gorilla/context
	requestedService     contextKey = iota // Used for gorilla/context
)

var (
	logger       *logrus.Entry
	config       environments.Config
	signingToken = os.Getenv("SIGNING_SECRET")
	bannedUsers  = map[string]string{ // Use a map for constant time lookups. Value doesn't matter
		// Hypothetically these should be universally unique, so we don't have to limit by env
		"d7b9dbed-9719-4856-5f19-23da2d0e3dec": "hidden",
	}
)

func init() {
	// Set the ENV environment variable to control dev/stage/prod behavior
	var err error
	config, err = environments.GetConfig(os.Getenv(environments.EnvVar))
	if err != nil {
		panic("Unable to determine configuration")
	}
	logger = config.GetLogger(ServiceName)
}

func main() {
	www.HttpServer()
}
