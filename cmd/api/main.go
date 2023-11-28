package main

import (
	"log"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/docs"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/di"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error at loading the env file using viper")
	}

	//	@title						Go + Gin Mobile-Mart
	//	@description				Mobile Mart - Your Ultimate Mobile Phone Shopping API 📱🛒
	//	@contact.name				API Support
	//	@securityDefinitions.apikey	BearerTokenAuth
	//	@in							header
	//	@name						Authorization
	//	@securityDefinitions.apikey	Refreshtoken
	//	@in							header
	//	@name						Refreshtoken
	//	@host						localhost:8080
	//	@BasePath					/
	//	@query.collection.format	multi

	docs.SwaggerInfo.Title = "Mobile_mart"
	// docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	// docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:7000"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	server, err := di.InitializeAPI(*config)
	if err != nil {
		log.Fatal("error for server creation")
	}

	server.Start()
}

//	@securityDefinitions.apikey	Refreshtoken
//	@in							header
//	@name						Refreshtoken
