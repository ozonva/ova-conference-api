package main

import (
	"github.com/rs/zerolog/log"
	server "ova-conference-api/internal/server/ova-conference-api"
	"ova-conference-api/internal/utils"
)

func main() {
	log.Print("Starting server...")
	config, err := utils.ReadConfigFromFile("test_config.json")()
	if err != nil {
		log.Err(err)
		log.Fatal()
	}
	log.Printf("Config is %v", config)

	if err := server.Start(config); err != nil {
		log.Err(err)
		log.Fatal()
	}
}
