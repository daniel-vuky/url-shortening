package main

import (
	"log"

	"github.com/daniel-vuky/url-shortening/internal/config"
	"github.com/daniel-vuky/url-shortening/internal/server"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	server.StartServer(config.Server.Port)
}
