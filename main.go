package main

import (
	"github.com/skyepic/privateapi/internal/infrastructure/web/webserver"
	"github.com/skyepic/privateapi/pkg/config"
	"github.com/skyepic/privateapi/pkg/database"
	"log"
)

func main() {
	config.LoadDotEnv()
	database.Connect()
	app := webserver.Setup()

	if err := app.ListenTLS(config.ServerUrl, config.CertFile, config.KeyFile); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
		log.Printf("Trying to run without SSL mode...")
		if err := app.Listen(config.ServerUrl); err != nil {
			log.Printf("Oops... Server is not running! Reason: %v", err)
		}
	}

}
