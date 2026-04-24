package main

import (
	"github.com/cassianobraz/GoGinCRUD/config"
	"github.com/cassianobraz/GoGinCRUD/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.NewLogger("main")
	// Initialize config
	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
