package main

import (
	"github.com/blackbarth/goopportunities.git/config"
	"github.com/blackbarth/goopportunities.git/router"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func main() {


	// Initialize the logger
	logger = config.GetLogger("main")

	// Initialize the configurations and dependencies
	if err := config.Init(); err != nil {
		logger.Errorf("Failed to initialize configurations: %v", err)
		return
	}
	defer config.CloseDB()

	// Initialize the router and start the server
	router.Initialize()

}