package main

import (
	"github.com/sw1pr0g/url_shortener/config"
	"github.com/sw1pr0g/url_shortener/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	} else {
		log.Printf("Config: %+v", cfg)
	}

	app.Run(cfg)
}
