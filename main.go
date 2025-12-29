package main

import (
	"github.com/WilliamOut/gopportunities/config"
	"github.com/WilliamOut/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initilization error: %v", err)
		return
	}

	//Initialize the router
	router.Initialize()
}
