package main

import (
	"os"
)

type serverConfig struct {
	port string
	addr string
}

// newServerConfig builds a serverConfig struct that is populated
// from environment variables or default values
func newServerConfig() *serverConfig {
	config := new(serverConfig)
	if config.addr = os.Getenv("HELLO_WORLD_ADDR"); config.addr == "" {
		config.addr = "0.0.0.0"
	}
	if config.port = os.Getenv("HELLO_WORLD_PORT"); config.port == "" {
		config.port = "8080"
	}
	return config
}
