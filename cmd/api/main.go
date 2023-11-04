package main

import (
	"log"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/di"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error at loading the env file using viper")
	}

	server, err := di.InitializeAPI(*config)
	if err != nil {
		log.Fatal("error for server creation")
	}

	server.Start()

}
