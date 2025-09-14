package main

import (
	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	// Run

	app.Run(cfg)
}
