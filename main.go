package main

import (
	"github.com/WilliamOut/gopportunities/config"
	"github.com/WilliamOut/gopportunities/handler"
	"github.com/WilliamOut/gopportunities/router"
)

var (
	logger *config.Logger
)

/*
{
  "role":"java",
  "company":"empresa x",
  "location":"salvador,BA",
  "link":"www.teste.com",
  "salary":3000
}
*/

// @title Gopportunities API
// @version 1.0
// @description API para gerenciamento de vagas de emprego.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	logger = config.GetLogger("main")
	//initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Erro na inicialização do config: %v", err)
		return
	}

	//Initialize handlers
	handler.InitHandler()

	//Initialize the router
	router.Initialize()
}
